package day04

import (
	"strconv"
	"strings"
)

func Overlaps(e1, e2 []int) bool {
	return e1[len(e1)-1] >= e2[0]
}

func FullyContains(e1, e2 []int) bool {
	return e1[0] <= e2[0] && e1[len(e1)-1] >= e2[len(e2)-1]
}

func Parse(input string) ([]int, []int, error) {
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
