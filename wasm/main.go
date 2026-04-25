package main

import (
	"github.com/mdfarhan20/connect-four/wasm/game"
)

const (
	Red = game.Red
	Yellow = game.Yellow
)

func main() {
	board := game.Board{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{Yellow, Yellow, Yellow, Yellow, 0, 0, 0},
		{Red, Red, Red, Yellow, Red, Red, 0},
	}

	g := game.CustomGameBoard(board)
	g.Winner()
	game.PrintBoard(*g)
}
