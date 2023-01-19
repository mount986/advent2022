package day17

import (
	"golang.org/x/exp/slices"
)

type Chamber struct {
	locations []Location
	Width     int
	Height    int
}

func NewChamber(width int) Chamber {
	var chamber Chamber
	chamber.Width = 7

	return chamber
}

func (c *Chamber) SettlePiece(s Shape) {
	for _, loc := range s.GetLocations() {
		c.AddLocation(loc)
	}

	height := s.LL.Y + s.Height()
	if height > c.Height {
		c.Height = height
	}
}

func (c Chamber) HasLocation(x, y int) bool {
	return slices.Contains(c.locations, NewLocation(x, y))
}

func (c *Chamber) AddLocation(loc Location) {
	c.locations = append(c.locations, loc)
}
