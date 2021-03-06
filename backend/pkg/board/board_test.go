package board

import (
	"math/rand"
	"testing"
	"time"
)

func TestBoardIsWon(t *testing.T) {
	tests := []string{
		"0011223",
		"0101010",
		"333333215444442122115620552",
		"444444326555553233226131666",
	}
	for _, test := range tests {
		board := boardFromString(test)
		if board.isWon() != RED {
			t.Errorf("Should be a winner %s", test)
		}
	}
}

func BenchmarkIsWon(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for n := 0; n < b.N; n++ {
		board := Board{}
		for board.isWon() == 0 {
			board.playMove(rand.Intn(7))
			if len(board.data) == 6*7 {
				break
			}
		}
	}
}
