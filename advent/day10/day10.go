package day10

import (
	"bufio"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func Run(scanner *bufio.Scanner) (int, error) {
	x := 1	
	sum := 0
	pause := false
	increase := 0
	var err error

	for cycle := 0; cycle < 220; cycle++ {
		if pause {
			x += increase
			pause = false
			continue
		}

		for scanner.Scan() {
			instruction := scanner.Text()

			if instruction[0:4] == "addx" {
				increase, err = strconv.Atoi(strings.Split(instruction, "\n")[1])
				if err != nil {
					return 0, err
				}
				pause = true
			}
		}

		if slices.Contains([]int{20, 60, 100, 140, 180, 220}, cycle) {
			sum +=  cycle * x
		}
	}
	
	return sum, nil
}