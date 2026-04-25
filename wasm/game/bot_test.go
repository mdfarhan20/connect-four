package game

import (
	"reflect"
	"testing"
)

func TestActions(t *testing.T) {
	t.Run("Empty board", func(t *testing.T) {
		game := NewGame()
		got := actions(*game)
		want := []int{0, 1, 2, 3, 4, 5, 6}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("actions() = %v, want %v", got, want)
		}
	})

	t.Run("Reachable full columns", func(t *testing.T) {
		game := NewGame()
		// Alternating moves to fill columns 0 and 6
		for i := 0; i < Height; i++ {
			game.MakeMove(0) // Red then Yellow then Red...
			game.MakeMove(6) // Yellow then Red then Yellow...
		}
		got := actions(*game)
		want := []int{1, 2, 3, 4, 5}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("actions() = %v, want %v", got, want)
		}
	})
}

func TestUtility(t *testing.T) {
	tests := []struct {
		name  string
		board Board
		want  int
	}{
		{
			name: "Red wins (Reachable)",
			board: Board{
				5: {Red, Yellow, Red, Yellow, Red, Yellow, Red},
				4: {Red, Yellow, Red, Yellow, 0, 0, 0},
				3: {Red, 0, 0, 0, 0, 0, 0},
				2: {Red, 0, 0, 0, 0, 0, 0},
			},
			want: 1,
		},
		{
			name: "Yellow wins (Reachable)",
			board: Board{
				5: {Red, Yellow, Red, Yellow, Red, Yellow},
				4: {Red, Yellow, Red, Yellow, 0, 0},
				3: {0, Yellow, 0, 0, 0, 0},
				2: {0, Yellow, 0, 0, 0, 0},
			},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := CustomGameBoard(tt.board)
			if got := utility(*game); got != tt.want {
				t.Errorf("utility() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestFindBestMove(t *testing.T) {
	t.Run("Immediate win with few options", func(t *testing.T) {
		// Board is almost full, only 2 spots left in row 5.
		// Red can win immediately in column 3.
		board := Board{
			{Red, Red, Red, Empty, Empty, Yellow, Yellow},
			{Red, Yellow, Red, Yellow, Red, Yellow, Red},
			{Yellow, Red, Yellow, Red, Yellow, Red, Yellow},
			{Red, Yellow, Red, Yellow, Red, Yellow, Red},
			{Yellow, Red, Yellow, Red, Yellow, Red, Yellow},
			{Red, Yellow, Red, Yellow, Red, Yellow, Red},
		}
		// 40 moves made (20 Red, 20 Yellow). Next is Red.
		game := CustomGameBoard(board)
		got := FindBestMove(*game)
		if got != 3 {
			t.Errorf("expected move 3 to win, got %d", got)
		}
	})

	t.Run("Must block with few options", func(t *testing.T) {
		// Board is almost full, only 2 spots left.
		// Yellow has 3 in a row, Red must block in column 3.
		board := Board{
			{Yellow, Yellow, Yellow, Empty, Empty, Red, Red},
			{Red, Yellow, Red, Yellow, Red, Yellow, Red},
			{Yellow, Red, Yellow, Red, Yellow, Red, Yellow},
			{Red, Yellow, Red, Yellow, Red, Yellow, Red},
			{Yellow, Red, Yellow, Red, Yellow, Red, Yellow},
			{Red, Yellow, Red, Yellow, Red, Yellow, Red},
		}
		// 40 moves made. Next is Red.
		game := CustomGameBoard(board)
		got := FindBestMove(*game)
		if got != 3 {
			t.Errorf("expected move 3 to block Yellow win, got %d", got)
		}
	})
}
