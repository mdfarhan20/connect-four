package game

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

		return explore(point.Move(direction))
 	}

	return explore(point)
}


func nextPlayer(c Cell) Cell {
	if c == Red {
		return Yellow
	}

	return Red
}