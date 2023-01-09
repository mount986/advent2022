package advent

import (
	"bufio"

	"github.com/mount986/advent2022/advent/day14"
)

func (a *Advent) Day14Part1() (int, error) {
	scanner := bufio.NewScanner(a.Input)	
	cave := day14.BuildMap(scanner)

	var start day14.Location
	start.X = 500
	start.Y = 0

	cave.FillWithSand(start, false)

	return cave.Sand, nil
}

func (a *Advent) Day14Part2() (int, error) {
	scanner := bufio.NewScanner(a.Input)	
	cave := day14.BuildMap(scanner)

	var start day14.Location
	start.X = 500
	start.Y = 0

	cave.FillWithSand(start, true)

	return cave.Sand, nil
}
