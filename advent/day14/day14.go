package day14

import (
	"bufio"
	"strings"
)

func BuildMap(scanner *bufio.Scanner) Map {
	var cave Map

	for scanner.Scan() {
		lines := scanner.Text()
		points := strings.Split(lines, " -> ")
	
		cave.SetCurrentLocation(
			NewLocation(
				strings.Split(points[0], ",")))
		
		for _, point := range points {
			cave.BuildTo(
				NewLocation(
					strings.Split(point, ",")))
		}

	}

	return cave
}