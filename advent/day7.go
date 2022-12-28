package advent

import (
	"bufio"
	"fmt"
	"regexp"
	"sort"
	"strconv"

	"github.com/mount986/advent2022/advent/day7"
)

func (a *Advent) Day7Part1() (int, error) {
	err := buildFileTree(a)
	if err != nil {
		return 0, err
	}

	_, sum := CollectSmall(day7.ROOT, 100_000)
	//calculate directory sizes
	//add it all up

	return sum, nil
}

func (a *Advent) Day7Part2() (int, error) {
	err := buildFileTree(a)
	if err != nil {
		return 0, err
	}

	systemSize := 70_000_000
	spaceNeeded := 30_000_000
	unusedSpace := systemSize - day7.ROOT.Size()

	largeDirs := CollectLarge(day7.ROOT, spaceNeeded-unusedSpace)

	sortf := func(i,j int) bool{ return largeDirs[i].Size() < largeDirs[j].Size()}

	sort.Slice(largeDirs, sortf)

	return largeDirs[0].Size(), nil
}

func buildFileTree(a *Advent) error {
	var cDir day7.Dir

	scanner := bufio.NewScanner(a.Input)
	for scanner.Scan() {
		line := scanner.Text()

		re, err := regexp.Compile(`(\$)? ?(\w+) ?(\S*)`)
		if err != nil {
			return nil
		}

		matches := re.FindStringSubmatch(line)

		if matches[1] == "$" {
			switch matches[2] {
			case "cd":
				cDir = cd(cDir, matches[3])
			case "ls":

			default:
				return fmt.Errorf("invalid command: %v", matches[2])
			}
		} else {
			if matches[2] != "dir" {
				if !cDir.ContainsFile(matches[3]) {
					size, err := strconv.Atoi(matches[2])
					if err != nil {
						return err
					}
					cDir.NewFile(matches[3], size)
				}
			}
		}
	}

	return nil
}

func cd(dir day7.Dir, sub string) day7.Dir {
	if sub == "/" {
		return day7.ROOT
	}

	if sub == ".." {
		return *dir.Parent
	}

	if dir.ContainsSub(sub) {
		return dir.Subs[sub]
	}

	return dir.NewSub(sub)
}

func CollectSmall(dir day7.Dir, max int) ([]day7.Dir, int) {
	var small []day7.Dir
	sum := 0

	if dir.Size() < max {
		small = append(small, dir)
		sum += dir.Size()
	}

	for _, sub := range dir.Subs {
		sm, su := CollectSmall(sub, max)
		small = append(small, sm...)
		sum += su
	}

	return small, sum
}

func CollectLarge(dir day7.Dir, min int) []day7.Dir {
	var largeDirs []day7.Dir

	if dir.Size() >= min {
		largeDirs = append(largeDirs, dir)

		for _, sub := range dir.Subs {
			largeDirs = append(largeDirs, CollectLarge(sub, min)...)
		}
	}

	return largeDirs
}
