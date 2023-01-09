package advent

import (
	"bufio"

	"github.com/mount986/advent2022/advent/day15"
)

func (a *Advent) Day15Part1() (int, error) {
	scanner := bufio.NewScanner(a.Input)	
	sensors := day15.BuildSensors(scanner)

	rowCount := day15.FindSensorsRow(sensors, 2_000_000)

	return rowCount, nil
}

func (a *Advent) Day15Part2() (int, error) {
	scanner := bufio.NewScanner(a.Input)	
	sensors := day15.BuildSensors(scanner)

	loc := day15.FindUnsensedLoc(sensors, 4_000_000)

	return (loc.X * 4_000_000 + loc.Y), nil
}
