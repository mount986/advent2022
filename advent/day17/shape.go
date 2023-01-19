package day17

import "log"

type Shape struct {
	Type    string
	LL      Location
	Settled bool
}

func NewShape(t string, height int) Shape {
	var shape Shape
	shape.Type = t
	shape.LL = NewLocation(2, height+3)
	shape.Settled = false

	return shape
}

func (s Shape) Width() int {
	switch s.Type {
	case "hLine":
		return 4
	case "cross", "angle":
		return 3
	case "square":
		return 2
	case "vLine":
		return 1
	}

	return 0
}

func (s Shape) Height() int {
	switch s.Type {
	case "hLine":
		return 1
	case "cross", "angle":
		return 3
	case "square":
		return 2
	case "vLine":
		return 4
	}

	return 0
}

func (s Shape) GetLocations() []Location {
	var locs []Location

	switch s.Type {
	case "hLine":
		locs = append(locs, NewLocation(s.LL.X, s.LL.Y))
		locs = append(locs, NewLocation(s.LL.X+1, s.LL.Y))
		locs = append(locs, NewLocation(s.LL.X+2, s.LL.Y))
		locs = append(locs, NewLocation(s.LL.X+3, s.LL.Y))
	case "cross":
		locs = append(locs, NewLocation(s.LL.X+1, s.LL.Y))
		locs = append(locs, NewLocation(s.LL.X, s.LL.Y+1))
		locs = append(locs, NewLocation(s.LL.X+1, s.LL.Y+1))
		locs = append(locs, NewLocation(s.LL.X+2, s.LL.Y+1))
		locs = append(locs, NewLocation(s.LL.X+1, s.LL.Y+2))
	case "angle":
		locs = append(locs, NewLocation(s.LL.X, s.LL.Y))
		locs = append(locs, NewLocation(s.LL.X+1, s.LL.Y))
		locs = append(locs, NewLocation(s.LL.X+2, s.LL.Y))
		locs = append(locs, NewLocation(s.LL.X+2, s.LL.Y+1))
		locs = append(locs, NewLocation(s.LL.X+2, s.LL.Y+2))
	case "vLine":
		locs = append(locs, NewLocation(s.LL.X, s.LL.Y))
		locs = append(locs, NewLocation(s.LL.X, s.LL.Y+1))
		locs = append(locs, NewLocation(s.LL.X, s.LL.Y+2))
		locs = append(locs, NewLocation(s.LL.X, s.LL.Y+3))
	case "square":
		locs = append(locs, NewLocation(s.LL.X, s.LL.Y))
		locs = append(locs, NewLocation(s.LL.X, s.LL.Y+1))
		locs = append(locs, NewLocation(s.LL.X+1, s.LL.Y))
		locs = append(locs, NewLocation(s.LL.X+1, s.LL.Y+1))
	}

	return locs
}

func (s *Shape) Move(chamber Chamber, wind byte) {
	s.MoveHorizontal(chamber, wind)
	s.MoveDown(chamber)
}

func (s *Shape) MoveHorizontal(chamber Chamber, wind byte) {
	var dir int
	switch string(wind) {
	case "<":
		if s.LL.X == 0 {
			return
		}
		dir = -1
	case ">":
		if s.LL.X == chamber.Width-s.Width() {
			return
		}
		dir = 1
	default:
		log.Fatalf("Unknown wind direction")
	}

	for _, loc := range s.GetLocations() {
		if chamber.HasLocation(loc.X+dir, loc.Y) {
			return
		}
	}

	s.LL.X += dir
}

func (s *Shape) MoveDown(chamber Chamber) {
	if s.LL.Y == 0 {
		s.Settled = true
		return
	}

	for _, loc := range s.GetLocations() {
		if chamber.HasLocation(loc.X, loc.Y-1) {
			s.Settled = true
			return
		}
	}

	s.LL.Y--
}
