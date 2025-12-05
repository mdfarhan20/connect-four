package game

type Cell uint8

// Board Cell Constants
const (
	Empty  Cell = 0
	Red    Cell = 1
	Yellow Cell = 2
)

// Board Dimensions
const (
	Height = 6
	Width  = 7
)

type Board [Height][Width]Cell

type Point struct {
	X, Y int
}

func (p Point) Move(d Direction) Point {
	offset := directionOffset[d]
	return Point{p.X + offset.X, p.Y + offset.Y}
}


func (p Point) isValid() bool {
	i, j := p.X, p.Y

	if i < 0 || j < 0 || i >= Height || j >= Width {
		return false
	}

	return true
}

type Direction uint8

const (
	Up Direction = iota
	Down
	Left
	Right
	RightUp
	RightDown
	LeftUp
	LeftDown
)

var directions = []Direction{Up, Down, Left, Right, RightUp, RightDown, LeftUp, LeftDown}

var directionOffset = [8]Point{
	Up:        {-1, 0},
	Down:      {1, 0},
	Left:      {0, -1},
	Right:     {0, 1},
	RightUp:   {-1, 1},
	RightDown: {1, 1},
	LeftUp:    {-1, -1},
	LeftDown:  {1, -1},
}