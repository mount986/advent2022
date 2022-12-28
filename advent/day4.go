package advent

import (
	"bufio"

	"github.com/mount986/advent2022/advent/day4"
)

func (a *Advent) Day4Part1() (int, error) {
	count := 0

	scanner := bufio.NewScanner(a.Input)
	for scanner.Scan() {
		line := scanner.Text()

		elf1, elf2, err := day4.Parse(line)
		if err != nil {
			return 0, err
		}

		if len(elf1) > len(elf2) {
			if day4.FullyContains(elf1, elf2) {
				count += 1
			}
		} else {
			if day4.FullyContains(elf2, elf1) {
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

		elf1, elf2, err := day4.Parse(line)
		if err != nil {
			return 0, err
		}

		if elf1[0] <= elf2[0] {
			if day4.Overlaps(elf1, elf2) {
				count += 1
			}
		} else {
			if day4.Overlaps(elf2, elf1) {
				count += 1
			}
		}
	}

	return count, nil
}
