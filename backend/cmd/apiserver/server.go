package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"regexp"
	"sync"
	"time"

	"golang.org/x/net/context"

	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/maxmcd/c4/backend/pkg/api"
	"github.com/maxmcd/c4/backend/pkg/board"
	"github.com/maxmcd/c4/backend/pkg/config"
	"github.com/maxmcd/c4/backend/pkg/models"
	"github.com/ttacon/libphonenumber"
	"google.golang.org/genproto/googleapis/rpc/code"
	google_rpc "google.golang.org/genproto/googleapis/rpc/status"
)

// Errors
var (
	ErrGameIdNotFound = func(id string) error { return fmt.Errorf("Game id (%s) not found", id) }
	ErrNotYourTurn    = func() error { return errors.New("Not your turn") }
	ErrGameIsOver     = func() error { return errors.New("Game is over") }
)

type APIServer struct {
	games         *GameStore
	backplane     *Backplane
	db            *gorm.DB
	doneChan      chan bool
	enterPoolChan chan PoolUser
}

func NewAPIServer() (*APIServer, error) {
	// TODO: Cleanly close database connection on API server close
	// TODO: Shutdown pool on API server close

	noDB := os.Getenv("NO_DB")
	db, err := gorm.Open("postgres", config.DatabaseUrl("postgres"))
	if err != nil && noDB == "" {
		return nil, err
	}

	enterPoolChan := make(chan PoolUser)
	doneChan := make(chan bool)
	games := GameStore{syncMap: sync.Map{}}
	go runPool(enterPoolChan, doneChan, &games)

	return &APIServer{
		games:         &games,
		backplane:     NewBackplane(),
		db:            db,
		enterPoolChan: enterPoolChan,
		doneChan:      doneChan,
	}, nil
}

func (s *APIServer) JoinPool(ctx context.Context, req *api.JoinPoolRequest) (*api.JoinPoolResponse, error) {
	userId := ctx.Value("userId").(int)
	var user models.User
	if err := s.db.Where("user_id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}
	gameChan := make(chan api.Game)
	poolUser := PoolUser{
		User:     user,
		Joined:   time.Now(),
		GameChan: gameChan,
	}
	s.enterPoolChan <- poolUser

	game := <-gameChan

	return &api.JoinPoolResponse{
		Game: &game,
	}, nil
}

func (s *APIServer) SendMove(ctx context.Context, req *api.SendMoveRequest) (*api.SendMoveResponse, error) {
	// Ensure that the game exists
	userId := ctx.Value("userId").(int)
	game, ok := s.games.Retrieve(req.Id)
	if !ok {
		return nil, ErrGameIdNotFound(req.Id)
	}
	command := req.Command
	if command == "resign" {
		userIdUint := uint(userId)
		game.Resigned = &userIdUint
	}

	// TODO: ensure this is threadsafe for duplicate moves from the same player

	if game.Resigned == nil {
		if (game.Board.WhoIsNext() == board.RED && uint(userId) != game.RedUser.ID) ||
			(game.Board.WhoIsNext() == board.BLACK && uint(userId) != game.BlackUser.ID) {
			return nil, ErrNotYourTurn()
		}
		isWon := game.Board.IsWon()
		if isWon != 0 {
			return nil, ErrGameIsOver()
		}
		err := game.Board.PlayMove(int(req.Column))
		if err != nil {
			return nil, err
		}
		if game.ClockStart == nil && len(game.Board.Data) == 2 {
			start := time.Now()
			game.ClockStart = &start
		}
		if isWon == 0 && game.Board.ClockShouldBeStarted() {
			if game.Board.WhoIsNext() == board.RED {
				game.BlackTimeRemaining = time.Now().Sub(*game.ClockStart) - (game.TimeControl - game.RedTimeRemaining)
			} else {
				game.RedTimeRemaining = time.Now().Sub(*game.ClockStart) - (game.TimeControl - game.BlackTimeRemaining)
			}
		}
	}

	s.games.Update(req.Id, game)
	s.backplane.Publish(fmt.Sprintf("game.%s", req.Id), &api.ReceiveMoveResponse{
		Command:            req.Command,
		Board:              game.Board.String(),
		BlackTimeRemaining: uint64(game.RedTimeRemaining),
		RedTimeRemaining:   uint64(game.BlackTimeRemaining),
	})

	return &api.SendMoveResponse{}, nil
}

func (s *APIServer) ReceiveMove(req *api.ReceiveMoveRequest, stream api.GameService_ReceiveMoveServer) error {
	glog.V(1).Infoln("ReceiveMove started...")
	defer glog.V(1).Infoln("ReceiveMove finished")
	if err := stream.Send(&api.ReceiveMoveResponse{
		Status: &google_rpc.Status{
			Code: int32(code.Code_OK),
		},
	}); err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(stream.Context())
	s.backplane.Subscribe(ctx, fmt.Sprintf("game.%s", req.Id), func(move interface{}) {
		glog.V(2).Infoln("ReceiveMove recevied module update: %+v", move)
		if err := stream.Send(move.(*api.ReceiveMoveResponse)); err != nil {
			// If we get an error, cancel the watch
			glog.Error(err)
			cancel()
		}
	})

	<-ctx.Done()

	return nil
}

func (s *APIServer) Authenticate(ctx context.Context, req *api.AuthenticateRequest) (*api.AuthenticateResponse, error) {
	glog.V(1).Infoln("Authenticate started...")
	defer glog.V(1).Infoln("Authenticate finished")

	num, err := libphonenumber.Parse(req.Phone, "US")
	if err != nil {
		return nil, err
	}
	user := models.User{
		Phone:       &sql.NullInt64{Int64: int64(num.GetNationalNumber()), Valid: true},
		CountryCode: &sql.NullInt64{Int64: int64(num.GetCountryCode()), Valid: true},
	}
	err = s.db.Where("phone = ? AND country_code = ?", user.Phone, user.CountryCode).First(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		if err := s.db.Create(&user).Error; err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	if !(user.Phone.Int64 == int64(2035559385) &&
		user.CountryCode.Int64 == int64(1)) {
		err = sendVerificationRequest(user)
		if err != nil {
			return nil, err
		}
	}

	return &api.AuthenticateResponse{}, nil
}

func (s *APIServer) SmsCodeCheck(ctx context.Context, req *api.SmsCodeCheckRequest) (*api.SmsCodeCheckResponse, error) {
	glog.V(1).Infoln("SmsCodeCheck started...")
	defer glog.V(1).Infoln("SmsCodeCheck finished")

	// TODO: format phone number
	num, err := libphonenumber.Parse(req.Phone, "US")
	if err != nil {
		return nil, err
	}

	user := models.User{}
	if err := s.db.Where("phone = ? AND country_code = ?",
		num.GetNationalNumber(), num.GetCountryCode()).First(&user).Error; err != nil {
		return nil, err
	}

	if user.Phone.Int64 != int64(2035559385) &&
		user.CountryCode.Int64 != int64(1) {
		if err := confirmVerificationRequest(user, req.Code); err != nil {
			return nil, err
		}
	}

	tokenString, err := user.Jwt()
	if err != nil {
		return nil, err
	}
	return &api.SmsCodeCheckResponse{
		Token:    tokenString,
		UserId:   uint64(user.ID),
		Username: user.UsernameString(),
	}, nil
}

// User Service Methods
func (s *APIServer) GetUser(ctx context.Context, req *api.GetUserRequest) (
	*api.GetUserResponse, error) {
	var user models.User
	if err := s.db.First(&user, ctx.Value("userId").(int)).Error; err != nil {
		return nil, err
	}

	return &api.GetUserResponse{
		Username: user.UsernameString(),
	}, nil
}

func (s *APIServer) UpdateUser(ctx context.Context, req *api.UpdateUserRequest) (
	*api.UpdateUserResponse, error) {
	// Username only for now
	if req.Username == "" {
		return nil, errors.New("Username can't be blank")
	}
	isValidChars := regexp.MustCompile(`^[0-9A-Za-z_]+$`).MatchString
	if !isValidChars(req.Username) {
		return nil, errors.New(
			"A username can only contain alphanumeric" +
				" characters (letters A-Z, numbers 0-9) with" +
				" the exception of underscores",
		)
	}

	var user models.User
	if err := s.db.First(&user, ctx.Value("userId").(int)).Error; err != nil {
		return nil, err
	}

	if err := s.db.Model(&user).Update("username", req.Username).Error; err != nil {
		return nil, err
	}

	return &api.UpdateUserResponse{}, nil
}
