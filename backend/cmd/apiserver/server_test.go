package main

import (
	"context"
	"os"
	"testing"

	"github.com/maxmcd/c4/backend/pkg/api"
	"github.com/maxmcd/c4/backend/pkg/models"
)

func TestBasicGame(t *testing.T) {
	os.Setenv("NO_DB", "truthy")
	apiserver, err := NewAPIServer()
	if err != nil {
		t.Fatal(err)
	}
	id := "foo"
	user := models.User{ID: uint(1)}
	user.FakeRating(1500 * 100)
	userContext := context.WithValue(
		context.TODO(), "userId", int(user.ID))
	user2 := models.User{ID: uint(2)}
	user2.FakeRating(1500 * 100)
	user2Context := context.WithValue(
		context.TODO(), "userId", int(user2.ID))
	_ = user2Context

	apiserver.games.Create(id, user, user2)
	_, err = apiserver.SendMove(userContext, &api.SendMoveRequest{
		Id:     id,
		Column: int32(1),
	})
	if err != nil {
		t.Error(err)
	}
	_, err = apiserver.SendMove(userContext, &api.SendMoveRequest{
		Id:     id,
		Column: int32(1),
	})
	if err == nil {
		t.Errorf("Should have failed")
	}

	for i := 0; i < 3; i++ {
		_, err = apiserver.SendMove(user2Context, &api.SendMoveRequest{
			Id:     id,
			Column: int32(2),
		})
		if err != nil {
			t.Error(err)
		}
		_, err = apiserver.SendMove(userContext, &api.SendMoveRequest{
			Id:     id,
			Column: int32(1),
		})
		if err != nil {
			t.Error(err)
		}
	}

	_, err = apiserver.SendMove(user2Context, &api.SendMoveRequest{
		Id:     id,
		Column: int32(2),
	})
	if err == nil {
		t.Errorf("game should be over at this point")
	}
	// game, _ := apiserver.games.Retrieve(id)
}
