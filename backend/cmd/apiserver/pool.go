package main

import (
	"math/rand"
	"sort"
	"time"

	"github.com/maxmcd/c4/backend/pkg/api"
	"github.com/maxmcd/c4/backend/pkg/models"
	"github.com/satori/go.uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	MATCH_CHECK_INTERVAL = time.Millisecond * 100
)

type PoolUser struct {
	User     models.User
	Joined   time.Time
	GameChan chan api.Game
}

type Pool []PoolUser

func (p Pool) Len() int           { return len(p) }
func (p Pool) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Pool) Less(i, j int) bool { return p[i].User.Rating() < p[j].User.Rating() }

func runPool(enterPoolChan chan PoolUser, done chan bool, gameStore *GameStore) {
	users := []PoolUser{}
	ticker := time.NewTicker(MATCH_CHECK_INTERVAL)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			matched := make(map[uint]bool)
			if len(users) == 1 {
				continue
			}
			for i := 0; i < len(users)-1; i++ {
				first := users[i]
				second := users[i+1]

				// for every second waited you're willing to add 100 points
				// to your rating spread
				ratingDiff := second.User.Rating() - first.User.Rating()
				firstAllowedSpread :=
					int(time.Now().Sub(first.Joined).Seconds() * 10000)
				secondAllowedSpread :=
					int(time.Now().Sub(second.Joined).Seconds() * 10000)
				// fmt.Println(ratingDiff, secondAllowedSpread, firstAllowedSpread)
				if ratingDiff-firstAllowedSpread < 0 && ratingDiff-secondAllowedSpread < 0 {
					i++ // skip the next round
					matched[first.User.ID] = true
					matched[second.User.ID] = true
					coinFlip := rand.Intn(1)
					var red models.User
					var black models.User
					if coinFlip == 0 {
						red = first.User
						black = second.User
					} else {
						red = second.User
						black = first.User
					}

					id := uuid.NewV4().String()
					gameStore.Create(id, red, black)
					game := api.Game{
						Id:        id,
						RedUser:   red.AsApiUser(),
						BlackUser: black.AsApiUser(),
					}

					first.GameChan <- game
					second.GameChan <- game
				}
			}
			users = func(users []PoolUser) []PoolUser {
				out := []PoolUser{}
				for _, user := range users {
					if !matched[user.User.ID] {
						out = append(out, user)
					}
				}
				return out
			}(users)
		case user := <-enterPoolChan:
			users = append(users, user)
			sort.Sort(Pool(users))
		}
	}
}

func findMatch(users []PoolUser) (red PoolUser, black PoolUser, err error) {
	return
}
