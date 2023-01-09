package day15

import (
	"bufio"
	"log"
	"regexp"
	"strconv"
)

type Sensor struct {
	Location Location
	Beacon Location
	distance int
}

func NewSensor(scanner *bufio.Scanner) Sensor {
	var s Sensor
	var err error
	
	re := regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)
	matches := re.FindStringSubmatch(scanner.Text())

	s.Location.X, err = strconv.Atoi(matches[1])
	if err != nil {
		log.Fatal(err)
	}
	
	s.Location.Y, err = strconv.Atoi(matches[2])
	if err != nil {
		log.Fatal(err)
	}
	
	s.Beacon.X, err = strconv.Atoi(matches[3])
	if err != nil {
		log.Fatal(err)
	}
	
	s.Beacon.Y, err = strconv.Atoi(matches[4])
	if err != nil {
		log.Fatal(err)
	}

	return s
}

func (s Sensor) Distance() int {
	if s.distance == 0 {
		s.distance = s.Location.ManhattanDistance(s.Beacon)
	}

	return s.distance
}