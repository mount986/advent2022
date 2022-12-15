package advent

import (
	"os"
	"testing"

	"github.com/mount986/advent2022/advent"
)

func TestDay2Part1(t *testing.T) {
	testFile := "data/day2"
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("could not open test data file %v", testFile)
	}
	defer file.Close()

	a := advent.Advent{
		Input: file,
	}

	score, err := a.Day2Part1()
	if err != nil {
		t.Fatal(err)
	}

	if score != 15 {
		t.Errorf("expected 15, got %v", score)
	}
}

func TestDay2Part2(t *testing.T) {
	testFile := "data/day2"
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("could not open test data file %v", testFile)
	}
	defer file.Close()

	a := advent.Advent{
		Input: file,
	}

	score, err := a.Day2Part2()
	if err != nil {
		t.Fatal(err)
	}

	if score != 12 {
		t.Errorf("expected 12, got %v", score)
	}
}
