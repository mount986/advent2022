package advent

import (
	"bufio"

	"github.com/mount986/advent2022/advent/day25"
)

func (a *Advent) Day25Part1old() (string, error) {
	scanner := bufio.NewScanner(a.Input)	
	var list []string

	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	sum := day25.SumSnafu(list)

	return sum, nil
}

func (a *Advent) Day25Part1() (string, error) {
	scanner := bufio.NewScanner(a.Input)	
	list := make(map[string]int)

	for scanner.Scan() {
		list[scanner.Text()] = day25.ToDec(scanner.Text())
	}

	sum := 0
	for _, num := range list {
		sum += num
	}

	return day25.ToSnafu(sum), nil
}

func (a *Advent) Day25Part2() (string, error) {
	scanner := bufio.NewScanner(a.Input)	
	var list []string

	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	sum := day25.SumSnafu(list)

	return sum, nil
}
