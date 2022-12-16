package advent

import (
	"bufio"
	"fmt"
	"strings"
)

func (a *Advent) Day3Part1() (int, error) {
	pri := 0

	scanner := bufio.NewScanner(a.Input)
	for scanner.Scan() {
		line := scanner.Text()

		comp1, comp2 := split(line)

		for _, ch := range comp1 {
			index := strings.Index(comp2, string(ch))
			if index != -1 {
				value, err := priority(ch)
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
				value, err := priority(ch)
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

func priority(ch rune) (int, error) {
	ascii := int(ch)

	switch {
	case ascii > 64 && ascii < 91:
		return ascii - 64 + 26, nil
	case ascii > 96 && ascii < 123:
		return ascii - 96, nil
	default:
		return -1, fmt.Errorf("not an alpha character: %v", ch)
	}

}

func split(bag string) (string, string) {
	return bag[:len(bag)/2], bag[len(bag)/2:]
}
