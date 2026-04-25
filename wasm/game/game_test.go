package game

import (
	"testing"
)

func TestNewGame(t *testing.T) {
	game := NewGame()
	if game.currentPlayer != Red {
		t.Errorf("expected current player Red, got %v", game.currentPlayer)
	}
	if game.moves != 0 {
		t.Errorf("expected 0 moves, got %d", game.moves)
	}
	if game.winner != Empty {
		t.Errorf("expected no winner, got %v", game.winner)
	}
}

func TestMakeMove(t *testing.T) {
	t.Run("Sequence of valid moves", func(t *testing.T) {
		game := NewGame()
		moves := []int{3, 3, 2, 4, 1, 5} // Alternating Red and Yellow
		for _, col := range moves {
			if err := game.MakeMove(col); err != nil {
				t.Fatalf("unexpected error at col %d: %v", col, err)
			}
		}
		if game.moves != 6 {
			t.Errorf("expected 6 moves, got %d", game.moves)
		}
		if game.currentPlayer != Red {
			t.Errorf("expected current player Red, got %v", game.currentPlayer)
		}
	})

	t.Run("Full column error", func(t *testing.T) {
		game := NewGame()
		for i := 0; i < Height; i++ {
			game.MakeMove(0)
		}
		err := game.MakeMove(0)
		if err == nil {
			t.Error("expected error for full column, got nil")
		}
	})
}

func TestWinner(t *testing.T) {
	tests := []struct {
		name  string
		board Board
		want  Cell
	}{
		{
			name: "Horizontal win Red (Supported)",
			board: Board{
				5: {Red, Red, Red, Red, Yellow, Yellow, Yellow},
			},
			want: Red,
		},
		{
			name: "Vertical win Yellow (Supported)",
			board: Board{
				5: {Red, Yellow, Red, 0, 0, 0, 0},
				4: {Red, Yellow, Red, 0, 0, 0, 0},
				3: {0, Yellow, Red, 0, 0, 0, 0},
				2: {0, Yellow, 0, 0, 0, 0, 0},
			},
			want: Yellow,
		},
		{
			name: "Diagonal Up-Right win Red (Supported)",
			board: Board{
				2: {0, 0, 0, Red},
				3: {0, 0, Red, Yellow},
				4: {0, Red, Yellow, Red},
				5: {Red, Yellow, Red, Yellow},
			},
			want: Red,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := CustomGameBoard(tt.board)
			if got := game.Winner(); got != tt.want {
				t.Errorf("Winner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDraw(t *testing.T) {
	t.Run("Realistic full board no winner", func(t *testing.T) {
		// A reachable draw pattern (Red starts):
		// R Y R Y R Y R
		// R Y R Y R Y R
		// Y R Y R Y R Y
		// Y R Y R Y R Y
		// R Y R Y R Y R
		// R Y R Y R Y R
		board := Board{
			{Red, Yellow, Red, Yellow, Red, Yellow, Red},
			{Red, Yellow, Red, Yellow, Red, Yellow, Red},
			{Yellow, Red, Yellow, Red, Yellow, Red, Yellow},
			{Yellow, Red, Yellow, Red, Yellow, Red, Yellow},
			{Red, Yellow, Red, Yellow, Red, Yellow, Red},
			{Red, Yellow, Red, Yellow, Red, Yellow, Red},
		}
		game := CustomGameBoard(board)
		if !game.isDraw() {
			t.Error("expected isDraw true for full board with no winner")
		}
	})
}
