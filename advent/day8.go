package advent

import (
	"bufio"
	"fmt"
	"strconv"
)

var xlen, ylen int

func (a *Advent) Day8Part1() (int, error) {
	grid, err := buildGrid(a)
	if err != nil {
		return 0, err
	}

	coords := scanVisible(grid)

	return len(uniqueStrings(coords)), nil
}

func (a *Advent) Day8Part2() (int, error) {
	grid, err := buildGrid(a)
	if err != nil {
		return 0, err
	}

	return scanScenic(grid), nil
}

func buildGrid(a *Advent) ([][]int, error) {
	grid := make([][]int, 0)

	scanner := bufio.NewScanner(a.Input)
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

func scanVisible(grid [][]int) []string {
	var coords []string

	xlen = len(grid[0])
	ylen = len(grid)

	coords = append(coords, scanLeftSide(grid)...)
	coords = append(coords, scanRightSide(grid)...)
	coords = append(coords, scanTop(grid)...)
	coords = append(coords, scanBottom(grid)...)

	return coords
}

func scanLeftSide(grid [][]int) []string {
	var coords []string

	for y := 0; y < ylen; y++ {
		coords = append(coords, scanRight(grid, -1, y)...)
	}

	return coords
}

func scanRightSide(grid [][]int) []string {
	var coords []string

	for y := 0; y < ylen; y++ {
		coords = append(coords, scanLeft(grid, xlen, y)...)
	}

	return coords
}

func scanTop(grid [][]int) []string {
	var coords []string

	for x := 0; x < xlen; x++ {
		coords = append(coords, scanDown(grid, x, -1)...)
	}

	return coords
}

func scanBottom(grid [][]int) []string {
	var coords []string

	for x := 0; x < xlen; x++ {
		coords = append(coords, scanUp(grid, x, ylen)...)
	}

	return coords
}

func scanScenic(grid [][]int) int {
	bestScore := 0

	xlen = len(grid[0])
	ylen = len(grid)

	for x := 0; x < xlen; x++ {
		for y := 0; y < ylen; y++ {
			left := len(scanLeft(grid, x, y))
			right := len(scanRight(grid, x, y))
			up := len(scanUp(grid, x, y))
			down := len(scanDown(grid, x, y))

			score := left * right * up * down

			if score > bestScore {
				bestScore = score
			}
		}
	}

	return bestScore
}

func scanLeft(grid [][]int, startX, startY int) []string {
	var coords []string
	tallest := -1
	startH := 0
	maxH:= 9

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
			break
		}
	}

	return coords
}

func scanRight(grid [][]int, startX, startY int) []string {
	var coords []string
	tallest := -1
	startH := 0
	maxH:= 9

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

func scanUp(grid [][]int, startX, startY int) []string {
	var coords []string
	tallest := -1
	startH := 0
	maxH:= 9

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

func scanDown(grid [][]int, startX, startY int) []string {
	var coords []string
	tallest := -1
	startH := 0
	maxH:= 9

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
