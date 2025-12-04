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
	Width = 7
)

type Board [Height][Width]Cell
