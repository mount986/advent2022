package advent

import (
	"bufio"
	"os"
	"testing"

	"github.com/mount986/advent2022/advent"
	"github.com/mount986/advent2022/advent/day15"
)

func TestDay15Part1(t *testing.T) {
	testFile := "data/day15"
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("could not open test data file %v", testFile)
	}
	defer file.Close()

	a := advent.Advent{
		Input: file,
	}

	scanner := bufio.NewScanner(a.Input)
	sensors := day15.BuildSensors(scanner)

	rowCount := day15.FindSensorsRow(sensors, 10)

	if rowCount != 26 {
		t.Errorf("expected 26, got %v", rowCount)
	}
}

func TestDay15Part2(t *testing.T) {
	testFile := "data/day15"
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("could not open test data file %v", testFile)
	}
	defer file.Close()

	a := advent.Advent{
		Input: file,
	}

	scanner := bufio.NewScanner(a.Input)
	sensors := day15.BuildSensors(scanner)

	loc := day15.FindUnsensedLoc(sensors, 20)
	if loc.X != 14 || loc.Y != 11 {
		t.Errorf("expected 14,11, got %v", loc)
	}
}

