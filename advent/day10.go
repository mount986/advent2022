package advent

import (
	"bufio"
	"strings"

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

func (a *Advent) Day10Part2() (string, error) {
	scanner := bufio.NewScanner(a.Input)
	output, err := day10.Display(scanner)
	if err != nil {
		return "", err
	}

	return "\n" + strings.Join(output, "\n"), nil
}
