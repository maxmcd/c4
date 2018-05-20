package main

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/maxmcd/c4/backend/pkg/api"
	"github.com/maxmcd/c4/backend/pkg/models"
)

func SkipTestRunPool(t *testing.T) {
	fakeRatings := []int{
		1500,
		2000,
		1200,
		1100,
		900,
		1300,
		800,
		700,
	}
	fakeUsers := []PoolUser{}
	for index, rating := range fakeRatings {
		user := models.User{ID: uint(index)}
		user.FakeRating(rating * 100)
		fakeUsers = append(fakeUsers, PoolUser{
			User:     user,
			Joined:   time.Now(),
			GameChan: make(chan api.Game, 100),
		})
	}
	start := time.Now()
	enterPoolChan := make(chan PoolUser, 20)
	done := make(chan bool)
	games := GameStore{syncMap: sync.Map{}}
	go runPool(enterPoolChan, done, &games)
	// wait for all to be matched
	for _, user := range fakeUsers {
		enterPoolChan <- user
	}
	fmt.Println("Waiting for matches")
	for _, user := range fakeUsers {
		<-user.GameChan
	}
	fmt.Println("all matched")

	games.syncMap.Range(func(key, value interface{}) bool {
		id := key.(string)
		game := value.(Game)
		fmt.Println(game.CreatedAt.Sub(start), id,
			game.BlackUser.Rating()/100, game.RedUser.Rating()/100)
		return true
	})
	close(done)
}
