package game

const DEPTH_LIMIT = 10

// Returns the best move for a game position
func FindBestMove(g Game) int {
	player := g.currentPlayer

	moves := actions(g)
	utils := make([]int, len(moves))

	var baseUtil int
	if player == Red {
		baseUtil = -2
	} else if player == Yellow {
		baseUtil = 2
	}

	for i, m := range moves {
		utils[i] = minimax(g, m, baseUtil, 0)
		if (player == Red && utils[i] > baseUtil) || (player == Yellow && utils[i] < baseUtil) {
			baseUtil = utils[i]
		}
	}

	idx := 0

	if player == Red {
		for i, util := range utils {
			if util > utils[idx] {
				idx = i
			}
		}
	} else {
		for i, util := range utils {
			if util < utils[idx] {
				idx = i
			}
		}
	}

	return moves[idx]
}

// Returns the utility value for a terminal game
func utility(g *Game) int {
	if g.winner == Red {
		return 1
	} else if g.winner == Yellow {
		return -1
	} else {
		return 0
	}
}

// Returns if a game has ended or not
func terminal(g *Game) bool {
	return g.isDraw() || g.Winner() != Empty
}

// Returns all possible moves
func actions(g Game) []int {
	moves := []int{}

	for i, count := range g.cellCount {
		if count < Height {
			moves = append(moves, i)
		}
	}

	return moves
}

// Applies minimax algorithm for a game and move
func minimax(game Game, move int, optimalUtil int, depth int) int {
	game.MakeMove(move)

	if terminal(&game) || depth >= DEPTH_LIMIT {
		return utility(&game)
	}

	var util int
	player := game.currentPlayer
	moves := actions(game)

	if player == Red {
		util = -2
	} else if player == Yellow {
		util = 2
	}

	for _, m := range moves {
		move_util := minimax(game, m, util, depth + 1)
		if player == Red {
			if optimalUtil <= move_util {
				return optimalUtil
			}
			util = max(util, move_util)
		} else {
			if optimalUtil >= move_util {
				return optimalUtil
			}
			util = min(util, move_util)
		}
	}

	return util
}
