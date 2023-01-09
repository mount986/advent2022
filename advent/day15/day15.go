package day15

import (
	"bufio"
	"math"
)

func BuildSensors(scanner *bufio.Scanner) []Sensor {
	var sensors []Sensor

	for scanner.Scan() {
		sensors = append(sensors, NewSensor(scanner))
	}

	return sensors
}

func FindSensorsRow(sensors []Sensor, row int) int {
	sensed := make(map[Location]bool)
	for _, sensor := range sensors {
		distY := int(math.Abs(float64(sensor.Location.Y - row)))

		rem := sensor.Distance() - distY
		if rem >= 0 {
			for x := sensor.Location.X - rem; x <= sensor.Location.X+rem; x++ {
				if !(sensor.Beacon.X == x && sensor.Beacon.Y == row) {
					sensed[NewLocation(x, row)] = true
				}
			}
		}
	}

	return len(sensed)
}

func FindUnsensedLoc(sensors []Sensor, max int) Location {
	for y := 0; y < max; y++ {
		col:
		for x := 0; x < max; x++ {
			for _, sensor := range sensors {
				if sensor.Location.ManhattanDistance(NewLocation(x, y)) <= sensor.Distance() {
					distY := int(math.Abs(float64(sensor.Location.Y - y)))
					distX := int(math.Abs(float64(sensor.Location.X - x)))
					remX := sensor.Distance() - distY
					x += remX + distX
					continue col
				}

			}

			return NewLocation(x, y)
		}
	}

	return NewLocation(0, 0)
}
