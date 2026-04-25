package game

import "testing"

func TestPointIsValid(t *testing.T) {
	tests := []struct {
		name  string
		point Point
		want  bool
	}{
		{"Top-left corner", Point{0, 0}, true},
		{"Top-right corner", Point{0, Width - 1}, true},
		{"Bottom-left corner", Point{Height - 1, 0}, true},
		{"Bottom-right corner", Point{Height - 1, Width - 1}, true},
		{"Interior point", Point{3, 3}, true},
		{"Outside top", Point{-1, 0}, false},
		{"Outside bottom", Point{Height, 0}, false},
		{"Outside left", Point{0, -1}, false},
		{"Outside right", Point{0, Width}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.point.isValid(); got != tt.want {
				t.Errorf("Point%v.isValid() = %v, want %v", tt.point, got, tt.want)
			}
		})
	}
}

func TestPointMove(t *testing.T) {
	start := Point{3, 3}
	tests := []struct {
		name      string
		direction Direction
		want      Point
	}{
		{"Move Up", Up, Point{2, 3}},
		{"Move Down", Down, Point{4, 3}},
		{"Move Left", Left, Point{3, 2}},
		{"Move Right", Right, Point{3, 4}},
		{"Move RightUp", RightUp, Point{2, 4}},
		{"Move RightDown", RightDown, Point{4, 4}},
		{"Move LeftUp", LeftUp, Point{2, 2}},
		{"Move LeftDown", LeftDown, Point{4, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := start.Move(tt.direction); got != tt.want {
				t.Errorf("Point%v.Move(%v) = %v, want %v", start, tt.direction, got, tt.want)
			}
		})
	}
}
