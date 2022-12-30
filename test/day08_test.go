package advent

import (
	"os"
	"testing"

	"github.com/mount986/advent2022/advent"
)

func TestDay8Part1(t *testing.T) {
	testFile := "data/day08"
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("could not open test data file %v", testFile)
	}
	defer file.Close()

	a := advent.Advent{
		Input: file,
	}

	v, err := a.Day8Part1()
	if err != nil {
		t.Fatal(err)
	}

	if v != 21 {
		t.Errorf("expected 21, got %v", v)
	}
}

func TestDay8Part2(t *testing.T) {
	testFile := "data/day08"
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("could not open test data file %v", testFile)
	}
	defer file.Close()

	a := advent.Advent{
		Input: file,
	}

	v, err := a.Day8Part2()
	if err != nil {
		t.Fatal(err)
	}

	if v != 8 {
		t.Errorf("expected 8, got %v", v)
	}
}
