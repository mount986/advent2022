package advent

import (
	"bufio"
	"sort"

	"github.com/mount986/advent2022/advent/day7"
)

func (a *Advent) Day7Part1() (int, error) {
	scanner := bufio.NewScanner(a.Input)
	err := day7.BuildFileTree(scanner)
	if err != nil {
		return 0, err
	}

	_, sum := day7.CollectSmall(day7.ROOT, 100_000)
	//calculate directory sizes
	//add it all up

	return sum, nil
}

func (a *Advent) Day7Part2() (int, error) {
	scanner := bufio.NewScanner(a.Input)
	err := day7.BuildFileTree(scanner)
	if err != nil {
		return 0, err
	}

	systemSize := 70_000_000
	spaceNeeded := 30_000_000
	unusedSpace := systemSize - day7.ROOT.Size()

	largeDirs := day7.CollectLarge(day7.ROOT, spaceNeeded-unusedSpace)

	sortf := func(i,j int) bool{ return largeDirs[i].Size() < largeDirs[j].Size()}

	sort.Slice(largeDirs, sortf)

	return largeDirs[0].Size(), nil
}

