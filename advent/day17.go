package advent

import (
	"bufio"

	"github.com/mount986/advent2022/advent/day17"
)

func (a *Advent) Day17Part1() (int, error) {
	scanner := bufio.NewScanner(a.Input)
	chamber := day17.RunSim(scanner, 2022)

	return chamber.Height, nil
}

func (a *Advent) Day17Part2() (int, error) {
	scanner := bufio.NewScanner(a.Input)
	chamber := day17.RunSim(scanner, 1_000_000_000_000)

	return chamber.Height, nil
}
