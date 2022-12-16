package advent

import (
	"os"
	"testing"

	"github.com/mount986/advent2022/advent"
)

func TestDay4Part1(t *testing.T) {
	testFile := "data/day4"
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("could not open test data file %v", testFile)
	}
	defer file.Close()

	a := advent.Advent{
		Input: file,
	}

	v, err := a.Day4Part1()
	if err != nil {
		t.Fatal(err)
	}

	if v != 2 {
		t.Errorf("expected 2, got %v", v)
	}
}

func TestDay4Part2(t *testing.T) {
	testFile := "data/day4"
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("could not open test data file %v", testFile)
	}
	defer file.Close()

	a := advent.Advent{
		Input: file,
	}

	v, err := a.Day4Part2()
	if err != nil {
		t.Fatal(err)
	}

	if v != 4 {
		t.Errorf("expected 4, got %v", v)
	}
}
