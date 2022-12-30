package day09

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var knots []map[string]int

func RunSim(scanner *bufio.Scanner, nKnots int) (map[string]int, error) {
	tracker := make(map[string]int)
	knots = make([]map[string]int, nKnots)

	for i := range knots {
		knots[i] = map[string]int{"x": 0, "y": 0}
	}

	for scanner.Scan() {
		move := strings.Split(scanner.Text(), " ")

		direction := move[0]
		distance, err := strconv.Atoi(move[1])
		if err != nil {
			return nil, err
		}

		for i := 0; i < distance; i++ {
			Moveknots(direction)
			MoveKnots(tracker)

		}

	}

	return tracker, nil
}

func Moveknots(direction string) {
	switch direction {
	case "U":
		knots[0]["y"]++
	case "D":
		knots[0]["y"]--
	case "L":
		knots[0]["x"]--
	case "R":
		knots[0]["x"]++
	}
}

func MoveKnots(tracker map[string]int) {
	for i := range knots {
		if i == 0 {
			continue
		}
		if math.Abs(float64(knots[i-1]["x"]-knots[i]["x"])) > 1 {
			if knots[i-1]["x"] > knots[i]["x"] {
				knots[i]["x"]++
			} else {
				knots[i]["x"]--
			}
			if math.Abs(float64(knots[i-1]["y"]-knots[i]["y"])) > 0 {
				if knots[i-1]["y"] > knots[i]["y"] {
					knots[i]["y"]++
				} else {
					knots[i]["y"]--
				}
			}
		} else if math.Abs(float64(knots[i-1]["y"]-knots[i]["y"])) > 1 {
			if knots[i-1]["y"] > knots[i]["y"] {
				knots[i]["y"]++
			} else {
				knots[i]["y"]--
			}
			if math.Abs(float64(knots[i-1]["x"]-knots[i]["x"])) > 0 {
				if knots[i-1]["x"] > knots[i]["x"] {
					knots[i]["x"]++
				} else {
					knots[i]["x"]--
				}
			}
		}

		if i == len(knots)-1 {
			tracker[fmt.Sprintf("%d,%d", knots[i]["x"], knots[i]["y"])] += 1
		}
	}
}
