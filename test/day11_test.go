package advent

import (
	"os"
	"testing"

	"github.com/mount986/advent2022/advent"
)

func TestDay11Part1(t *testing.T) {
	testFile := "data/day11"
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("could not open test data file %v", testFile)
	}
	defer file.Close()

	a := advent.Advent{
		Input: file,
	}

	v, err := a.Day11Part1()
	if err != nil {
		t.Fatal(err)
	}

	if v != 10605 {
		t.Errorf("expected 10605, got %v", v)
	}
}

func TestDay11Part2(t *testing.T) {
	testFile := "data/day11"
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("could not open test data file %v", testFile)
	}
	defer file.Close()

	a := advent.Advent{
		Input: file,
	}

	v, err := a.Day11Part2()
	if err != nil {
		t.Fatal(err)
	}

	if v != 2713310158 {
		t.Errorf("expected 2713310158, got %v", v)
	}
}
