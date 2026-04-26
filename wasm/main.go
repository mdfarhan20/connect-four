package main

import (
	"fmt"

	"github.com/mdfarhan20/connect-four/wasm/game"
)

const (
	Red    = game.Red
	Yellow = game.Yellow
)

func main() {
	board := game.Board{
    {0,      0,      0,      0,      0,      0,      0     }, // Row 0
    {0,      0,      0,      0,      0,      0,      0     }, // Row 1
    {0,      0,      0,      0,      0,      0,      0     }, // Row 2
    {Red,    Yellow, Red,    0,      Yellow, Red,    Yellow}, // Row 3
    {Yellow, Red,    Yellow, Red,    Red,    Yellow, Red   }, // Row 4
    {Red,    Yellow, Red,    Yellow, Yellow, Red,    Yellow}, // Row 5
}

	g := game.CustomGameBoard(board)
	game.PrintBoard(*g)
	fmt.Println("Best move", game.FindBestMove(*g))
}
