package advent

import (
	"bufio"

	"github.com/mount986/advent2022/advent/day13"
)

func (a *Advent) Day13Part1() (int, error) {
	scanner := bufio.NewScanner(a.Input)
	pList, err := day13.BuildPackages(scanner)
	if err != nil {
		return 0, err
	}

	var valid []int
	for index, p := range pList {
		if p.ValidPackage() {
			valid = append(valid, index+1)
		}
	}

	sum := 0
	for _, v := range valid {
		sum += v
	}

	return sum, nil
}

func (a *Advent) Day13Part2() (int, error) {
	return 0, nil
}
