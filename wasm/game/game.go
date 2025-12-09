package game

import "fmt"

type Game struct {
	board Board
	currentPlayer Cell
	winner Cell
	moves uint8
	cellCount [Width]int
	maxCellHeight int
}


func NewGame() *Game {
	return &Game{
		currentPlayer: Red, 
		winner: Empty,
		moves: 0,
		cellCount: [Width]int{},
		maxCellHeight: 0,
	}
}


func (g Game) GetBoard() Board {
	return g.board
} 


func (g *Game) MakeMove(col int) error {
	if g.cellCount[col] >= Height {
		return fmt.Errorf("column is full")
	}

	row := g.cellCount[col]
	g.board[row][col] = g.currentPlayer
	g.currentPlayer = nextPlayer(g.currentPlayer)

	g.cellCount[col]++
	g.moves++

	if g.cellCount[col] > g.maxCellHeight {
		g.maxCellHeight = g.cellCount[col]
	}

	return nil
}


func (g *Game) Winner() Cell {
	if g.moves < 8 {
		return 0
	}

	for i := 0; i < Height; i++ {
		for j := 0; j < Width; j++ {
			if g.board[i][j] == Empty {
				continue
			}

			directions = getPossibleDirections(g.board, Point{i, j})

			for _, d := range directions {
				if gameWon := exploreDirection(g.board, Point{i, j}, d); gameWon {
					g.winner = g.board[i][j]
					return g.winner
				}
			}
		}
	}

	return Empty
}


func (g Game) isDraw() bool {
	return g.moves == (Height * Width)
}
