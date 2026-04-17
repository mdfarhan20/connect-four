package game


// Creates a custom game from a board
func CustomGameBoard(b Board) *Game {
	game := NewGame()

	game.board = b

	for _, row := range game.board {
		for i, cell := range row {
			if cell != Empty {
				// Updating column cell count and maxCellHeight
				game.cellCount[i]++
				if game.maxCellHeight < game.cellCount[i] {
					game.maxCellHeight = game.cellCount[i]
				}

				// Updating total game moves
				game.moves++
			}
		}
	}

	// Updating the current player
	if game.moves % 2 == 0 {
		game.currentPlayer = Red
	} else {
		game.currentPlayer = Yellow
	}

	// Checking for game winner
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

		// If new position is invalid or the cell in that point is not the same
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


// Returns the total moves made in a board
// func getTotalMoves(b Board) int {
// 	var moves int

// 	for _, row := range b {
// 		for _, cell := range row {
// 			if cell != Empty {
// 				moves++
// 			}
// 		} 
// 	}

// 	return moves
// }


// Returns the current player to move from a given board
// func currentPlayer(b Board) Cell {
// 	moves := getTotalMoves(b)

// 	if moves % 2 == 0 {
// 		return Red
// 	} 

// 	return Yellow
// }