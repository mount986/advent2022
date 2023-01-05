package advent

import (
	"bufio"

	"github.com/mount986/advent2022/advent/day11"
)

func (a *Advent) Day11Part1() (int, error) {
	scanner := bufio.NewScanner(a.Input)
	monkeys, err := day11.Init(scanner)
	if err != nil {
		return 0, err
	}
	v, err := day11.Run(monkeys, 20, true)
	if err != nil {
		return 0, err
	}

	return v, nil
}

func (a *Advent) Day11Part2() (int, error) {
	scanner := bufio.NewScanner(a.Input)
	monkeys, err := day11.Init(scanner)
	if err != nil {
		return 0, err
	}
	v, err := day11.Run(monkeys, 10_000, false)
	if err != nil {
		return 0, err
	}

	return v, nil
}
