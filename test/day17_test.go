package advent

import (
	"os"
	"testing"

	"github.com/mount986/advent2022/advent"
)

func TestDay17Part1(t *testing.T) {
	testFile := "data/day17"
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("could not open test data file %v", testFile)
	}
	defer file.Close()

	a := advent.Advent{
		Input: file,
	}

	v, err := a.Day17Part1()
	if err != nil {
		t.Fatal(err)
	}

	if v != 3068 {
		t.Errorf("expected 3068, got %v", v)
	}
}

func TestDay17Part2(t *testing.T) {
	testFile := "data/day17"
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("could not open test data file %v", testFile)
	}
	defer file.Close()

	a := advent.Advent{
		Input: file,
	}

	v, err := a.Day17Part2()
	if err != nil {
		t.Fatal(err)
	}

	if v != 1_514_285_714_288 {
		t.Errorf("expected 1,514,285,714,288, got %v", v)
	}
}