package advent

import (
	"os"
	"testing"

	"github.com/mount986/advent2022/advent"
)

func TestDay5Part1(t *testing.T) {
	testFile := "data/day5"
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("could not open test data file %v", testFile)
	}
	defer file.Close()

	a := advent.Advent{
		Input: file,
	}

	v, err := a.Day5Part1()
	if err != nil {
		t.Fatal(err)
	}

	if v != "CMZ" {
		t.Errorf("expected CMZ, got %v", v)
	}
}

func TestDay5Part2(t *testing.T) {
	testFile := "data/day5"
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("could not open test data file %v", testFile)
	}
	defer file.Close()

	a := advent.Advent{
		Input: file,
	}

	v, err := a.Day5Part2()
	if err != nil {
		t.Fatal(err)
	}

	if v != "MCD" {
		t.Errorf("expected MCD, got %v", v)
	}
}
