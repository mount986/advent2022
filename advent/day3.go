package advent

import (
	"bufio"
	"strings"

	"github.com/mount986/advent2022/advent/day3"
)

func (a *Advent) Day3Part1() (int, error) {
	pri := 0

	scanner := bufio.NewScanner(a.Input)
	for scanner.Scan() {
		line := scanner.Text()

		comp1, comp2 := day3.Split(line)

		for _, ch := range comp1 {
			index := strings.Index(comp2, string(ch))
			if index != -1 {
				value, err := day3.Priority(ch)
				if err != nil {
					return 0, err
				}

				pri += value
				break
			}
		}
	}

	return pri, nil
}

func (a *Advent) Day3Part2() (int, error) {
	pri := 0

	scanner := bufio.NewScanner(a.Input)
	for scanner.Scan() {
		line1 := scanner.Text()
		scanner.Scan()
		line2 := scanner.Text()
		scanner.Scan()
		line3 := scanner.Text()

		var runes []rune

		for _, ch := range line1 {
			index := strings.Index(line2, string(ch))
			if index != -1 {
				runes = append(runes, ch)
			}
		}

		for _, ch := range string(runes) {
			index := strings.Index(line3, string(ch))
			if index != -1 {
				value, err := day3.Priority(ch)
				if err != nil {
					return 0, err
				}

				pri += value
				break
			}
		}

	}

	return pri, nil
}

