package game

import "fmt"


// Creates a custom game from a board
func CustomGameBoard(b Board) *Game {
	game := NewGame()

	game.board = b

	for _, row := range game.board {
		for i, cell := range row {
			if cell != Empty {
				game.cellCount[i]++
				if game.maxCellHeight < game.cellCount[i] {
					game.maxCellHeight = game.cellCount[i]
				}

				game.moves++
			}
		}
	}

	if game.moves % 2 == 0 {
		game.currentPlayer = Red
	} else {
		game.currentPlayer = Yellow
	}

	game.Winner()

	return game
}


// Returns all possible directions to explore for a win
func getPossibleDirections(board Board, point Point) []Direction {
	i, j := point.X, point.Y
	cell := board[i][j]
	
	var possibleDirections []Direction
	for _, direction := range directions {
		newPosition := point.Move(direction)

		if !newPosition.isValid() || board[newPosition.X][newPosition.Y] != cell {
			continue
		}

		possibleDirections = append(possibleDirections, direction)
	}

	return possibleDirections
}


// Explores a particular direction for a win
func exploreDirection(board Board, point Point, direction Direction) bool {
	cell := board[point.X][point.Y]
	count := 0

	var explore func (Point) bool
	explore = func (p Point) bool {
		if !p.isValid() || board[p.X][p.Y] != cell {
			return false
		}

		count++
		if count == 4 {
			return true
		}

		return explore(p.Move(direction))
 	}

	return explore(point)
}


// Returns the next player based on current player
func nextPlayer(c Cell) Cell {
	if c == Red {
		return Yellow
	}

	return Red
}


const (
	reset  = "\033[0m"
	red    = "\033[31m"
	yellow = "\033[33m"
)

func cellToSymbol(c Cell) string {
	switch c {
	case Red:
		return red + "●" + reset
	case Yellow:
		return yellow + "●" + reset
	default:
		return "·"
	}
}


func PrintBoard(g Game) {
	board := g.board

	fmt.Println("  0   1   2   3   4   5   6")
	fmt.Println("┌───┬───┬───┬───┬───┬───┬───┐")

	for i := 0; i < Height; i++ {
		fmt.Print("│")
		for j := 0; j < Width; j++ {
			fmt.Print(" " + cellToSymbol(board[i][j]) + " │")
		}
		fmt.Println()

		if i < Height-1 {
			fmt.Println("├───┼───┼───┼───┼───┼───┼───┤")
		}
	}

	fmt.Println("└───┴───┴───┴───┴───┴───┴───┘")

	if g.winner != Empty {
		fmt.Printf("Winner: %s\n", cellToSymbol(g.winner))
	} else if g.isDraw() {
		fmt.Println("Draw!")
	} else {
		fmt.Printf("Current player: %s\n", cellToSymbol(g.currentPlayer))
	}
}
