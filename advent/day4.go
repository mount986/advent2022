package advent

import (
	"bufio"
	"strconv"
	"strings"
)

func (a *Advent) Day4Part1() (int, error) {
	count := 0

	scanner := bufio.NewScanner(a.Input)
	for scanner.Scan() {
		line := scanner.Text()

		elf1, elf2, err := parse(line)
		if err != nil {
			return 0, err
		}

		if len(elf1) > len(elf2) {
			if fullyContains(elf1, elf2) {
				count += 1
			}
		} else {
			if fullyContains(elf2, elf1) {
				count += 1
			}
		}
	}

	return count, nil
}

func (a *Advent) Day4Part2() (int, error) {
	count := 0

	scanner := bufio.NewScanner(a.Input)
	for scanner.Scan() {
		line := scanner.Text()

		elf1, elf2, err := parse(line)
		if err != nil {
			return 0, err
		}

		if elf1[0] <= elf2[0] {
			if overlaps(elf1, elf2) {
				count += 1
			}
		} else {
			if overlaps(elf2, elf1) {
				count += 1
			}
		}
	}

	return count, nil
}

func overlaps(e1, e2 []int) bool {
	return e1[len(e1)-1] >= e2[0]
}

func fullyContains(e1, e2 []int) bool {
	return e1[0] <= e2[0] && e1[len(e1)-1] >= e2[len(e2)-1]
}

func parse(input string) ([]int, []int, error) {
	ranges := strings.Split(input, ",")
	var elves [][]int

	for _, r := range ranges {
		var areas []int

		bounds := strings.Split(r, "-")
		start, err := strconv.Atoi(bounds[0])
		if err != nil {
			return nil, nil, err
		}

		end, err := strconv.Atoi(bounds[1])
		if err != nil {
			return nil, nil, err
		}
		for i := start; i <= end; i++ {
			areas = append(areas, i)
		}

		elves = append(elves, areas)
	}

	return elves[0], elves[1], nil
}
