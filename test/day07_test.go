package advent

import (
	"os"
	"testing"

	"github.com/mount986/advent2022/advent"
)

func TestDay7Part1(t *testing.T) {
	testFile := "data/day07"
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("could not open test data file %v", testFile)
	}
	defer file.Close()

	a := advent.Advent{
		Input: file,
	}

	v, err := a.Day7Part1()
	if err != nil {
		t.Fatal(err)
	}

	if v != 95437 {
		t.Errorf("expected 95437, got %v", v)
	}
}

func TestDay7Part2(t *testing.T) {
	testFile := "data/day07"
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("could not open test data file %v", testFile)
	}
	defer file.Close()

	a := advent.Advent{
		Input: file,
	}

	v, err := a.Day7Part2()
	if err != nil {
		t.Fatal(err)
	}

	if v != 24933642 {
		t.Errorf("expected 24933642, got %v", v)
	}
}
