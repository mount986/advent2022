package day15

import (
	"math"
)

type Location struct {
	X, Y int
}

func NewLocation(x, y int) Location {
	return Location{X: x, Y: y}
}

func (l Location) ManhattanDistance(o Location) int {
	return int(math.Abs(float64(l.X - o.X)) + math.Abs(float64(l.Y - o.Y)))
}
