package main

import (
	"sync"
	"time"

	"github.com/maxmcd/c4/backend/pkg/models"
)

type Game struct {
	Id                 string
	RedUser            models.User
	BlackUser          models.User
	CreatedAt          time.Time
	RedTimeRemaining   time.Duration
	BlackTimeRemaining time.Duration
}

type GameStore struct {
	syncMap sync.Map
}

func (g *GameStore) Create(id string, redUser models.User, blackUser models.User) {
	g.syncMap.Store(id, Game{
		Id:                 id,
		RedUser:            redUser,
		BlackUser:          blackUser,
		CreatedAt:          time.Now(),
		RedTimeRemaining:   time.Minute,
		BlackTimeRemaining: time.Minute,
	})
}

func (g *GameStore) Retrieve(id string) (Game, bool) {
	gameInt, ok := g.syncMap.Load(id)
	var game Game
	if ok {
		game = gameInt.(Game)
	}
	return game, ok
}

func (g *GameStore) Update(id string, game Game) {
	g.syncMap.Store(id, game)
}
