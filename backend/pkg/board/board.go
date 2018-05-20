package board

import (
	"errors"
)

const (
	RED          = 1
	BLACK        = 2
	BOARD_WIDTH  = 7
	BOARD_HEIGHT = 6
)

var runeIntMap = map[rune]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
}
var intStringMap = map[int]string{
	0: "0",
	1: "1",
	2: "2",
	3: "3",
	4: "4",
	5: "5",
	6: "6",
}

type Board struct {
	Data []int
}

func BoardFromString(moves string) (board Board) {
	for _, runeValue := range moves {
		board.Data = append(board.Data, runeIntMap[runeValue])
	}
	return
}

func (b *Board) String() (out string) {
	for _, column := range b.Data {
		out += intStringMap[column]
	}
	return
}

func (b *Board) PlayMove(column int) error {
	if len(b.Data) == 6*7 {
		return errors.New("Board full")
	}
	var count int
	for _, move := range b.Data {
		if move == column {
			count += 1
		}
	}
	if count == 6 {
		return errors.New("Column full")
	}
	b.Data = append(b.Data, column)
	return nil
}

func (b *Board) WhoIsNext() int {
	if len(b.Data)%2 == 1 {
		return BLACK
	}
	return RED
}

func (b *Board) ClockShouldBeStarted() bool {
	return len(b.Data) > 1
}

func (b *Board) asMatrix() (board [7][6]int) {
	var counters [7]int
	for index, column := range b.Data {
		board[column][counters[column]] = (index % 2) + 1
		counters[column] += 1
	}
	return
}

func (b *Board) IsWon() int {
	board := b.asMatrix()
	// horizontalCheck
	for j := 0; j < BOARD_HEIGHT-3; j++ {
		for i := 0; i < BOARD_WIDTH; i++ {
			if (board[i][j] == RED || board[i][j] == BLACK) &&
				board[i][j] == board[i][j+1] &&
				board[i][j] == board[i][j+2] &&
				board[i][j] == board[i][j+3] {
				return board[i][j]
			}
		}
	}
	// verticalCheck
	for i := 0; i < BOARD_WIDTH-3; i++ {
		for j := 0; j < BOARD_HEIGHT; j++ {
			if (board[i][j] == RED || board[i][j] == BLACK) &&
				board[i][j] == board[i+1][j] &&
				board[i][j] == board[i+2][j] &&
				board[i][j] == board[i+3][j] {
				return board[i][j]
			}
		}
	}
	// ascendingDiagonalCheck
	for i := 3; i < BOARD_WIDTH; i++ {
		for j := 0; j < BOARD_HEIGHT-3; j++ {
			if (board[i][j] == RED || board[i][j] == BLACK) &&
				board[i][j] == board[i-1][j+1] &&
				board[i][j] == board[i-2][j+2] &&
				board[i][j] == board[i-3][j+3] {
				return board[i][j]
			}
		}
	}
	// descendingDiagonalCheck
	for i := 3; i < BOARD_WIDTH; i++ {
		for j := 3; j < BOARD_HEIGHT; j++ {
			if (board[i][j] == RED || board[i][j] == BLACK) &&
				board[i][j] == board[i-1][j-1] &&
				board[i][j] == board[i-2][j-2] &&
				board[i][j] == board[i-3][j-3] {
				return board[i][j]
			}
		}
	}
	return 0
}
