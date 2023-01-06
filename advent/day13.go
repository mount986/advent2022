package advent

import (
	"bufio"
	"sort"
	"strings"

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
	scanner := bufio.NewScanner(a.Input)
	pList, err := day13.BuildPackages(scanner)
	if err != nil {
		return 0, err
	}

	dividers := 
`[[2]]
[[6]]`

	scanner = bufio.NewScanner(strings.NewReader(dividers))
	dList, err := day13.BuildPackages(scanner)
	if err != nil {
		return 0, err
	}

	pList = append(pList, dList[0])

	lists := day13.CombineLists(pList)

	sortf := func(i, j int) bool { return lists[i].Compare(lists[j]) < 0 }

	sort.Slice(lists, sortf)

	var twoIndex, sixIndex int

	for index := range lists {
		if lists[index].ToString() == "[[2]]" {
			twoIndex = index + 1
		}
		if lists[index].ToString() == "[[6]]" {
			sixIndex = index + 1
		}
	}

	return twoIndex * sixIndex, nil
}
