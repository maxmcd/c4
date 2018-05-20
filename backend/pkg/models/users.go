package models

import (
	"database/sql"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/maxmcd/c4/backend/pkg/api"
	"github.com/maxmcd/c4/backend/pkg/config"
)

// User wraps user entries in the database
type User struct {
	ID             uint `gorm:"primary_key"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Phone          *sql.NullInt64
	CountryCode    *sql.NullInt64
	Username       *sql.NullString
	GamesPlayed    int
	RatingScratch  int
	NexmoRequestID string
}

func (u *User) FakeRating(desiredRating int) {
	u.RatingScratch = desiredRating
	u.GamesPlayed = 1
}

func (u *User) UpdateRating(won bool, opponentRating int) {
	u.RatingScratch += opponentRating
	if won {
		u.RatingScratch += 40000 // 400.00
	} else {
		u.RatingScratch -= 40000 // -400.00
	}
}

func (u *User) Rating() int {
	if u.GamesPlayed == 0 {
		return 150000 // 1500.00
	}
	return (u.RatingScratch) / u.GamesPlayed
}

func (u *User) UsernameString() (username string) {
	if u.Username != nil {
		return u.Username.String
	}
	return
}

func (u *User) Jwt() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// NOTE: This claim may cause issue due to clock drift (esp in VMs)
		"nbf": time.Now().Unix(),
		// json doesn't sucessfully deserialize large uints
		"user_id": fmt.Sprintf("%d", u.ID),
	})
	tokenString, err := token.SignedString([]byte(config.Auth.Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (u *User) PhoneNumber() string {
	if u.Phone != nil && u.CountryCode != nil {
		// TODO probably want to pad correct number of zeros?
		return fmt.Sprintf("%d%d", u.CountryCode.Int64, u.Phone.Int64)
	}
	return ""
}

func (u *User) AsApiUser() *api.User {
	return &api.User{
		Id:       uint64(u.ID),
		Username: u.UsernameString(),
	}
}
