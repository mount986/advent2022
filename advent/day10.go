package advent

import (
	"bufio"

	"github.com/mount986/advent2022/advent/day10"
)

func (a *Advent) Day10Part1() (int, error) {
	scanner := bufio.NewScanner(a.Input)
	sum, err := day10.Run(scanner)
	if err != nil {
		return 0, err
	}

	return sum, nil
}

func (a *Advent) Day10Part2() (int, error) {
	scanner := bufio.NewScanner(a.Input)
	sum, err := day10.Run(scanner)
	if err != nil {
		return 0, err
	}

	return sum, nil
}
