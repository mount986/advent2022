package advent

import (
	"bufio"

	"github.com/mount986/advent2022/advent/day8"
)


func (a *Advent) Day8Part1() (int, error) {
	scanner := bufio.NewScanner(a.Input)
	grid, err := day8.BuildGrid(scanner)
	if err != nil {
		return 0, err
	}

	coords := day8.ScanVisible(grid)

	return len(uniqueStrings(coords)), nil
}

func (a *Advent) Day8Part2() (int, error) {
	scanner := bufio.NewScanner(a.Input)
	grid, err := day8.BuildGrid(scanner)
	if err != nil {
		return 0, err
	}

	return day8.ScanScenic(grid), nil
}