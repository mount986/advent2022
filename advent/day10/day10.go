package day10

import (
	"bufio"
	"strconv"

	"golang.org/x/exp/slices"
)

func Run(scanner *bufio.Scanner) (int, error) {
	x := 1
	sum := 0
	scanner.Split(bufio.ScanWords)

	for cycle := 1; cycle <= 220; cycle++ {
		if slices.Contains([]int{20, 60, 100, 140, 180, 220}, cycle) {
			sum += cycle * x
		}

		scanner.Scan()
		instruction := scanner.Text()
		switch instruction {
		case "addx":
		case "noop":
		default:
			value, err := strconv.Atoi(instruction)
			if err != nil {
				return 0, err
			}
			x += value
		}
	}

	return sum, nil
}

func Display(scanner *bufio.Scanner) ([]string, error) {
	output := make([]string, 6)

	scanner.Split(bufio.ScanWords)
	sprite := 1
	cycle := 0
	row := -1

	for scanner.Scan() {
		if cycle == 240 {
			break
		}

		x := cycle %40
		if x % 40 == 0 {
			row++
		}

		if sprite >= x - 1 && sprite <= x + 1 {
			output[row] += "#"
		} else {
			output[row] += "."
		}
		
		instruction := scanner.Text()
		switch instruction {
		case "addx":
		case "noop":
		default:
			value, err := strconv.Atoi(instruction)
			if err != nil {
				return output, err
			}
			sprite += value
		}
		cycle++
	}

	return output, nil
}
