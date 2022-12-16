package advent

import (
	"os"
	"testing"
	"github.com/mount986/advent2022/advent"
)

func TestDay1Part1(t *testing.T) {
	testFile := "data/day1"
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("could not open test data file %v", testFile)
	}
	defer file.Close()

	a := advent.Advent{
		Input: file,
	}

	max, err := a.Day1Part1()
	if err != nil {
		t.Fatal(err)
	}

	if max != 24000 {
		t.Errorf("expected 2400, got %v", max)
	}
}

func TestDay1Part2(t *testing.T) {
	testFile := "data/day1"
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("could not open test data file %v", testFile)
	}
	defer file.Close()

	a := advent.Advent{
		Input: file,
	}
	
	max, err := a.Day1Part2()
	if err != nil {
		t.Fatal(err)
	}

	if max != 45000 {
		t.Errorf("expected 45000, got %v", max)
	}
}
