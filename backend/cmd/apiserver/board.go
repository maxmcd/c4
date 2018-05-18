package main

import (
	"errors"
	"fmt"
	"strings"
)

const (
	InitialBoard = ""
	RED          = "r"
	BLACK        = "b"
)

type Board struct {
	data string
}

var runeIntMap map[rune]int

func init() {
	runeIntMap = map[rune]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
	}
}

func (b *Board) playMove(column int) error {
	if len(b.data) == 6*7 {
		return errors.New("Board full")
	}
	if strings.Count(b.data, fmt.Sprintf("%d", column)) == 6 {
		return errors.New("Column full")
	}
	b.data = fmt.Sprintf("%s%d", b.data, column)
	return nil
}

func (b *Board) whoIsNext() string {
	if len(b.data)%2 == 1 {
		return BLACK
	}
	return RED
}

func (b *Board) asMatrix() (board [7][7]rune) {
	var counters [7]int
	colors := [2]rune{'r', 'b'}
	for index, runeValue := range b.data {
		column := runeIntMap[runeValue]
		board[column][counters[column]] = colors[index%2]
		counters[column] += 1
	}
	return
}

func (b *Board) isWon() rune {
	board := b.asMatrix()
	// horizontalCheck
	for j := 0; j < 7-3; j++ {
		for i := 0; i < 7; i++ {
			if (board[i][j] == 'r' || board[i][j] == 'b') &&
				board[i][j] == board[i][j+1] &&
				board[i][j] == board[i][j+2] &&
				board[i][j] == board[i][j+3] {
				return board[i][j]
			}
		}
	}
	// verticalCheck
	for i := 0; i < 7-3; i++ {
		for j := 0; j < 7; j++ {
			if (board[i][j] == 'r' || board[i][j] == 'b') &&
				board[i][j] == board[i+1][j] &&
				board[i][j] == board[i+2][j] &&
				board[i][j] == board[i+3][j] {
				return board[i][j]
			}
		}
	}
	// ascendingDiagonalCheck
	for i := 3; i < 7; i++ {
		for j := 0; j < 7-3; j++ {
			if (board[i][j] == 'r' || board[i][j] == 'b') &&
				board[i][j] == board[i-1][j+1] &&
				board[i][j] == board[i-2][j+2] &&
				board[i][j] == board[i-3][j+3] {
				return board[i][j]
			}
		}
	}
	// descendingDiagonalCheck
	for i := 3; i < 7; i++ {
		for j := 3; j < 7; j++ {
			if (board[i][j] == 'r' || board[i][j] == 'b') &&
				board[i][j] == board[i-1][j-1] &&
				board[i][j] == board[i-2][j-2] &&
				board[i][j] == board[i-3][j-3] {
				return board[i][j]
			}
		}
	}
	var zeroRune rune
	return zeroRune
}
