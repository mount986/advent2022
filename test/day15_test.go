package advent

import (
	"os"
	"testing"

	"github.com/mount986/advent2022/advent"
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

	v, err := a.Day15Part1()
	if err != nil {
		t.Fatal(err)
	}

	if v != 24 {
		t.Errorf("expected 24, got %v", v)
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

	v, err := a.Day15Part2()
	if err != nil {
		t.Fatal(err)
	}

	if v != 93 {
		t.Errorf("expected 93, got %v", v)
	}
}