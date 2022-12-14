package advent

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Day2Part1(input *os.File) (int, error) {
	win := map[string]string{
		"A": "Y",
		"B": "Z",
		"C": "X",
	}
	tie := map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}
	lose := map[string]string{
		"A": "Z",
		"B": "X",
		"C": "Y",
	}
	points := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	score := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		moves := strings.Split(line, " ")

		score += points[moves[1]]
		if win[moves[0]] == moves[1] {
			score += 6
		} else if tie[moves[0]] == moves[1] {
			score += 3
		} else if lose[moves[0]] != moves[1] {
			return 0, fmt.Errorf("Bad input: %v %v", moves[0], moves[1])
		}
	}

	return score, nil
}

func Day2Part2(input *os.File) (int, error) {
	strategy := map[string]string{
		"X": "lose",
		"Y": "tie",
		"Z": "win",
	}

	winPoints := map[string]int{
		"A": 2,
		"B": 3,
		"C": 1,
	}

	tiePoints := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	losePoints := map[string]int{
		"A": 3,
		"B": 1,
		"C": 2,
	}

	score := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		moves := strings.Split(line, " ")

		if strategy[moves[1]] == "win" {
			score += 6
			score += winPoints[moves[0]]
		} else if strategy[moves[1]] == "tie" {
			score += 3
			score += tiePoints[moves[0]]
		} else if strategy[moves[1]] == "lose" {
			score += losePoints[moves[0]]
		} else {
			return 0, fmt.Errorf("Unexpected input: %v %v", moves[0], moves[1])
		}
	}

	return score, nil
}
