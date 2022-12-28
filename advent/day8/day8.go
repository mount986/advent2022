package day8

import (
	"bufio"
	"fmt"
	"strconv"
)

var xlen, ylen int

func BuildGrid(scanner *bufio.Scanner) ([][]int, error) {
	grid := make([][]int, 0)

	scanner.Split(bufio.ScanRunes)

	i := 0
	grid = append(grid, [][]int{make([]int, 0)}...)
	for scanner.Scan() {
		r := scanner.Text()

		if r == "\n" {
			i++
			grid = append(grid, [][]int{make([]int, 0)}...)
		} else {
			t, err := strconv.Atoi(r)
			if err != nil {
				return grid, err
			}
			grid[i] = append(grid[i], t)
		}
	}

	return grid, nil
}

func ScanVisible(grid [][]int) []string {
	var coords []string

	xlen = len(grid[0])
	ylen = len(grid)

	coords = append(coords, ScanLeftSide(grid)...)
	coords = append(coords, ScanRightSide(grid)...)
	coords = append(coords, ScanTop(grid)...)
	coords = append(coords, ScanBottom(grid)...)

	return coords
}

func ScanLeftSide(grid [][]int) []string {
	var coords []string

	for y := 0; y < ylen; y++ {
		coords = append(coords, ScanRight(grid, -1, y)...)
	}

	return coords
}

func ScanRightSide(grid [][]int) []string {
	var coords []string

	for y := 0; y < ylen; y++ {
		coords = append(coords, ScanLeft(grid, xlen, y)...)
	}

	return coords
}

func ScanTop(grid [][]int) []string {
	var coords []string

	for x := 0; x < xlen; x++ {
		coords = append(coords, ScanDown(grid, x, -1)...)
	}

	return coords
}

func ScanBottom(grid [][]int) []string {
	var coords []string

	for x := 0; x < xlen; x++ {
		coords = append(coords, ScanUp(grid, x, ylen)...)
	}

	return coords
}

func ScanScenic(grid [][]int) int {
	bestScore := 0

	xlen = len(grid[0])
	ylen = len(grid)

	for x := 0; x < xlen; x++ {
		for y := 0; y < ylen; y++ {
			left := len(ScanLeft(grid, x, y))
			right := len(ScanRight(grid, x, y))
			up := len(ScanUp(grid, x, y))
			down := len(ScanDown(grid, x, y))

			score := left * right * up * down

			if score > bestScore {
				bestScore = score
			}
		}
	}

	return bestScore
}

func ScanLeft(grid [][]int, startX, startY int) []string {
	var coords []string
	tallest := -1
	startH := 0
	maxH := 9

	if startX >= 0 && startX < xlen {
		startH = grid[startY][startX]
		maxH = startH
	}

	for x := startX - 1; x >= 0; x-- {
		if grid[startY][x] < startH || grid[startY][x] > tallest {
			coords = append(coords, fmt.Sprintf("%d,%d", x, startY))
			tallest = grid[startY][x]
		}

		if tallest >= maxH {
			if tallest >= maxH {
				break
			}
		}
	}

	return coords
}

func ScanRight(grid [][]int, startX, startY int) []string {
	var coords []string
	tallest := -1
	startH := 0
	maxH := 9

	if startX >= 0 && startX < xlen {
		startH = grid[startY][startX]
		maxH = startH
	}

	for x := startX + 1; x < xlen; x++ {
		if grid[startY][x] < startH || grid[startY][x] > tallest {
			coords = append(coords, fmt.Sprintf("%d,%d", x, startY))
			tallest = grid[startY][x]
		}

		if tallest >= maxH {
			break
		}
	}

	return coords
}

func ScanUp(grid [][]int, startX, startY int) []string {
	var coords []string
	tallest := -1
	startH := 0
	maxH := 9

	if startY >= 0 && startY < xlen {
		startH = grid[startY][startX]
		maxH = startH
	}

	for y := startY - 1; y >= 0; y-- {
		if grid[y][startX] < startH || grid[y][startX] > tallest {
			coords = append(coords, fmt.Sprintf("%d,%d", startX, y))
			tallest = grid[y][startX]
		}

		if tallest >= maxH {
			break
		}
	}

	return coords
}

func ScanDown(grid [][]int, startX, startY int) []string {
	var coords []string
	tallest := -1
	startH := 0
	maxH := 9

	if startY >= 0 && startY < xlen {
		startH = grid[startY][startX]
		maxH = startH
	}

	for y := startY + 1; y < ylen; y++ {
		if grid[y][startX] < startH || grid[y][startX] > tallest {
			coords = append(coords, fmt.Sprintf("%d,%d", startX, y))
			tallest = grid[y][startX]
		}

		if tallest >= maxH {
			break
		}
	}

	return coords
}
