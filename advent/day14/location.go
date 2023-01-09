package day14

import (
	"log"
	"strconv"
)

type Location struct {
	X, Y int
}

func NewLocation(coords []string) Location {
	x, err := strconv.Atoi(coords[0])
	if err != nil {
		log.Fatal(err)
	}

	y, err := strconv.Atoi(coords[1])
	if err != nil {
		log.Fatal(err)
	}

	return Location{X: x, Y: y}
}
