package advent

import (
	"bufio"

	"github.com/mount986/advent2022/advent/day08"
)

func (a *Advent) Day8Part1() (int, error) {
	scanner := bufio.NewScanner(a.Input)
	grid, err := day08.BuildGrid(scanner)
	if err != nil {
		return 0, err
	}

	coords := day08.ScanVisible(grid)

	return len(uniqueStrings(coords)), nil
}

func (a *Advent) Day8Part2() (int, error) {
	scanner := bufio.NewScanner(a.Input)
	grid, err := day08.BuildGrid(scanner)
	if err != nil {
		return 0, err
	}

	return day08.ScanScenic(grid), nil
}
