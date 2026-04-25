package game

import (
	"reflect"
	"testing"
)

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
			if got := nextPlayer(tt.input); got != tt.want {
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
		wantCellCount [Width]int
		wantMaxHeight int
	}{
		{
			name:          "Empty board",
			board:         Board{},
			wantMoves:     0,
			wantCurrent:   Red,
			wantCellCount: [Width]int{0, 0, 0, 0, 0, 0, 0},
			wantMaxHeight: 0,
		},
		{
			name: "Valid 3 moves (Red, Yellow, Red)",
			board: Board{
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{Red, Yellow, Red, 0, 0, 0, 0},
			},
			wantMoves:     3,
			wantCurrent:   Yellow,
			wantCellCount: [Width]int{1, 1, 1, 0, 0, 0, 0},
			wantMaxHeight: 1,
		},
		{
			name: "Valid 4 moves (Red, Yellow, Red, Yellow)",
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
			wantCellCount: [Width]int{1, 1, 1, 1, 0, 0, 0},
			wantMaxHeight: 1,
		},
		{
			name: "Stacked pieces (Red, Yellow alternating)",
			board: Board{
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{Yellow, 0, 0, 0, 0, 0, 0}, // Move 4
				{Red, 0, 0, 0, 0, 0, 0},    // Move 3
				{Yellow, 0, 0, 0, 0, 0, 0}, // Move 2
				{Red, 0, 0, 0, 0, 0, 0},    // Move 1
			},
			wantMoves:     4,
			wantCurrent:   Red,
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
			if !reflect.DeepEqual(game.cellCount, tt.wantCellCount) {
				t.Errorf("cellCount = %v, want %v", game.cellCount, tt.wantCellCount)
			}
			if game.maxCellHeight != tt.wantMaxHeight {
				t.Errorf("maxCellHeight = %v, want %v", game.maxCellHeight, tt.wantMaxHeight)
			}
		})
	}
}

func TestGetPossibleDirections(t *testing.T) {
	// A reachable board fragment
	board := Board{
		5: {Red, Yellow, Red, Yellow, 0, 0, 0},
		4: {Red, Yellow, Red, Yellow, 0, 0, 0},
	}

	t.Run("Boundary of stack", func(t *testing.T) {
		got := getPossibleDirections(board, Point{5, 0})
		// Red at (5,0) has:
		// Right: Yellow (Fail)
		// Up: Red (Success)
		// RightUp: Yellow (Fail)
		if len(got) != 1 || got[0] != Up {
			t.Errorf("got %v directions, want [Up]", got)
		}
	})
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
			name: "Horizontal 4 (Bottom row)",
			board: Board{
				5: {Red, Red, Red, Red, Yellow, Yellow, Yellow},
			},
			point:     Point{5, 0},
			direction: Right,
			want:      true,
		},
		{
			name: "Vertical 4 (Column 0)",
			board: Board{
				5: {Red, Yellow},
				4: {Red, Yellow},
				3: {Red, Yellow},
				2: {Red, Yellow},
			},
			point:     Point{5, 0},
			direction: Up,
			want:      true,
		},
		{
			name: "Diagonal Up-Right 4 (Supported)",
			board: Board{
				2: {0, 0, 0, Red},
				3: {0, 0, Red, Yellow},
				4: {0, Red, Yellow, Red},
				5: {Red, Yellow, Red, Yellow},
			},
			point:     Point{5, 0},
			direction: RightUp,
			want:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exploreDirection(tt.board, tt.point, tt.direction); got != tt.want {
				t.Errorf("exploreDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}
