package advent

import (
	"bufio"

	"github.com/mount986/advent2022/advent/day09"
)

func (a *Advent) Day9Part1() (int, error) {
	scanner := bufio.NewScanner(a.Input)
	tracker, err := day09.RunSim(scanner, 2)
	if err != nil {
		return 0, err
	}

	return len(tracker), nil
}

func (a *Advent) Day9Part2() (int, error) {
	scanner := bufio.NewScanner(a.Input)
	tracker, err := day09.RunSim(scanner, 10)
	if err != nil {
		return 0, err
	}

	return len(tracker), nil
}
