package game

import "testing"

func TestNextPlayer(t *testing.T) {
	tests := []struct {
		name  string
		input Cell
		want  Cell
	}{
		{"Red to Yellow", Red, Yellow},
		{"Yellow to Red", Yellow, Red},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := nextPlayer(tt.input)
			if got != tt.want {
				t.Errorf("nextPlayer(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestCustomGameBoard(t *testing.T) {
	tests := []struct {
		name          string
		board         Board
		wantMoves     uint8
		wantCurrent   Cell
		wantWinner    Cell
		wantCellCount [Width]int
		wantMaxHeight int
	}{
		{
			name:          "Empty board",
			board:         Board{},
			wantMoves:     0,
			wantCurrent:   Red,
			wantWinner:    Empty,
			wantCellCount: [Width]int{0, 0, 0, 0, 0, 0, 0},
			wantMaxHeight: 0,
		},
		{
			name: "Single Red move",
			board: Board{
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{Red, 0, 0, 0, 0, 0, 0},
			},
			wantMoves:     1,
			wantCurrent:   Yellow,
			wantWinner:    Empty,
			wantCellCount: [Width]int{1, 0, 0, 0, 0, 0, 0},
			wantMaxHeight: 1,
		},
		{
			name: "Two moves",
			board: Board{
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{Red, Yellow, 0, 0, 0, 0, 0},
			},
			wantMoves:     2,
			wantCurrent:   Red,
			wantWinner:    Empty,
			wantCellCount: [Width]int{1, 1, 0, 0, 0, 0, 0},
			wantMaxHeight: 1,
		},
		{
			name: "Multiple pieces",
			board: Board{
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{Red, Yellow, Red, Yellow, 0, 0, 0},
			},
			wantMoves:     4,
			wantCurrent:   Red,
			wantWinner:    Empty,
			wantCellCount: [Width]int{1, 1, 1, 1, 0, 0, 0},
			wantMaxHeight: 1,
		},
		{
			name: "Max cell height",
			board: Board{
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{Red, 0, 0, 0, 0, 0, 0},
				{Yellow, 0, 0, 0, 0, 0, 0},
				{Red, 0, 0, 0, 0, 0, 0},
				{Yellow, 0, 0, 0, 0, 0, 0},
			},
			wantMoves:     4,
			wantCurrent:   Red,
			wantWinner:    Empty,
			wantCellCount: [Width]int{4, 0, 0, 0, 0, 0, 0},
			wantMaxHeight: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := CustomGameBoard(tt.board)

			if game.moves != tt.wantMoves {
				t.Errorf("moves = %v, want %v", game.moves, tt.wantMoves)
			}
			if game.currentPlayer != tt.wantCurrent {
				t.Errorf("currentPlayer = %v, want %v", game.currentPlayer, tt.wantCurrent)
			}
			if game.winner != tt.wantWinner {
				t.Errorf("winner = %v, want %v", game.winner, tt.wantWinner)
			}
			for col, want := range tt.wantCellCount {
				if game.cellCount[col] != want {
					t.Errorf("cellCount[%d] = %v, want %v", col, game.cellCount[col], want)
				}
			}
			if game.maxCellHeight != tt.wantMaxHeight {
				t.Errorf("maxCellHeight = %v, want %v", game.maxCellHeight, tt.wantMaxHeight)
			}
		})
	}
}

func TestGetPossibleDirections(t *testing.T) {
	tests := []struct {
		name      string
		board     Board
		point     Point
		wantCount int
		wantDirs  []Direction
	}{
		{
			name: "Isolated cell",
			board: Board{
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, Red, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
			},
			point:     Point{3, 3},
			wantCount: 0,
			wantDirs:  nil,
		},
		{
			name: "Single neighbor right",
			board: Board{
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, Red, Red, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
			},
			point:     Point{3, 3},
			wantCount: 1,
			wantDirs:  []Direction{Right},
		},
		{
			name: "Multiple neighbors",
			board: Board{
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, Red, 0, 0, 0, 0},
				{0, 0, 0, Red, Red, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
			},
			point:     Point{3, 3},
			wantCount: 2,
			wantDirs:  []Direction{Right, LeftUp},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dirs := getPossibleDirections(tt.board, tt.point)

			if len(dirs) != tt.wantCount {
				t.Errorf("getPossibleDirections count = %d, want %d", len(dirs), tt.wantCount)
			}
			if tt.wantDirs != nil {
				for i, want := range tt.wantDirs {
					if i >= len(dirs) || dirs[i] != want {
						t.Errorf("GetPossibleDirections = %v, want %v", dirs, tt.wantDirs)
						break
					}
				}
			}
		})
	}
}

func TestExploreDirection(t *testing.T) {
	tests := []struct {
		name      string
		board     Board
		point     Point
		direction Direction
		want      bool
	}{
		{
			name: "Three in a row",
			board: Board{
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, Red, Red, Red, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
			},
			point:     Point{3, 2},
			direction: Right,
			want:      false,
		},
		{
			name: "Four in a row",
			board: Board{
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, Red, Red, Red, Red, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
			},
			point:     Point{3, 1},
			direction: Right,
			want:      true,
		},
		{
			name: "Five in a row",
			board: Board{
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{Red, Red, Red, Red, Red, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
			},
			point:     Point{3, 0},
			direction: Right,
			want:      true,
		},
		{
			name: "Different colors",
			board: Board{
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, Red, Yellow, Red, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
			},
			point:     Point{3, 2},
			direction: Right,
			want:      false,
		},
		{
			name: "Break in chain",
			board: Board{
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, Red, Red, 0, Red, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
			},
			point:     Point{3, 2},
			direction: Right,
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := exploreDirection(tt.board, tt.point, tt.direction)
			if got != tt.want {
				t.Errorf("exploreDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}