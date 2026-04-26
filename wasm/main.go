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
    {Red,    Red,    Red,    0,      Yellow, Yellow, Yellow}, // Row 1
    {Yellow, Yellow, Red,    Red,    Yellow, Yellow, Red   }, // Row 2
    {Red,    Red,    Yellow, Yellow, Red,    Red,    Yellow}, // Row 3
    {Yellow, Yellow, Red,    Red,    Yellow, Yellow, Red   }, // Row 4
    {Red,    Red,    Yellow, Yellow, Red,    Red,    Yellow}, // Row 5
}

	g := game.CustomGameBoard(board)
	game.PrintBoard(*g)
	fmt.Println("Best move", game.FindBestMove(*g))
}
