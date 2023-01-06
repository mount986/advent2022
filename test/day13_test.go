package advent

import (
	"os"
	"testing"

	"github.com/mount986/advent2022/advent"
)

func TestDay13Part1(t *testing.T) {
	testFile := "data/day13"
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("could not open test data file %v", testFile)
	}
	defer file.Close()

	a := advent.Advent{
		Input: file,
	}

	v, err := a.Day13Part1()
	if err != nil {
		t.Fatal(err)
	}

	if v != 13 {
		t.Errorf("expected 13, got %v", v)
	}
}