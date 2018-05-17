package main

import (
	"database/sql"
	"errors"
	"fmt"

	"golang.org/x/net/context"

	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/maxmcd/c4/backend/pkg/api"
	"github.com/maxmcd/c4/backend/pkg/config"
	"github.com/maxmcd/c4/backend/pkg/models"
	"github.com/ttacon/libphonenumber"
	"google.golang.org/genproto/googleapis/rpc/code"
	google_rpc "google.golang.org/genproto/googleapis/rpc/status"
)

// Errors
var (
	ErrGameIdNotFound = func(id string) error { return fmt.Errorf("Game id (%s) not found", id) }
)

type APIServer struct {
	// TODO: Persist these to some type of data store later
	// map of stream id to stream object
	games map[string]*api.Game
	// TODO: These updates need to be streamed over a real backplane
	backplane *Backplane
	db        *gorm.DB
}

func NewAPIServer() (*APIServer, error) {
	// TODO: Cleanly close database connection on API server close
	db, err := gorm.Open("postgres", config.DatabaseUrl("postgres"))
	if err != nil {
		return nil, err
	}
	return &APIServer{
		games:     make(map[string]*api.Game),
		backplane: NewBackplane(),
		db:        db,
	}, nil
}

func (s *APIServer) JoinPool(ctx context.Context, req *api.JoinPoolRequest) (*api.JoinPoolResponse, error) {
	return &api.JoinPoolResponse{}, nil
}

func (s *APIServer) SendMove(ctx context.Context, req *api.SendMoveRequest) (*api.SendMoveResponse, error) {
	// Ensure that the game exists
	if _, ok := s.games[req.Id]; !ok {
		return nil, ErrGameIdNotFound(req.Id)
	}

	// Publish this module to anyone listening in on this game id
	s.backplane.Publish(fmt.Sprintf("game.%s", req.Id), req)

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
		moveRequest := move.(*api.SendMoveRequest)
		if err := stream.Send(&api.ReceiveMoveResponse{
			Command: moveRequest.Command,
			Board:   moveRequest.Board,
		}); err != nil {
			// If we get an error, cancel the watch
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

	var request_id string
	if user.Phone.Int64 == int64(2035559385) &&
		user.CountryCode.Int64 == int64(1) {
		request_id = "testing"
	} else {
		request_id, err = sendVerificationRequest(user)
		if err != nil {
			return nil, err
		}
	}

	if err := s.db.Model(&user).Update("nexmo_request_id", request_id).Error; err != nil {
		return nil, err
	}

	return &api.AuthenticateResponse{
		RequestId: request_id,
	}, nil
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

	if req.RequestId == "" || req.RequestId != user.NexmoRequestID {
		return nil, errors.New("Invalid request id")
	}

	if user.Phone.Int64 != int64(2035559385) &&
		user.CountryCode.Int64 != int64(1) {
		if err := confirmVerificationRequest(req.RequestId, req.Code); err != nil {
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
