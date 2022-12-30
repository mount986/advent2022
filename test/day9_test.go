package advent

import (
	"os"
	"testing"

	"github.com/mount986/advent2022/advent"
)

func TestDay9Part1(t *testing.T) {
	testFile := "data/day9a"
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("could not open test data file %v", testFile)
	}
	defer file.Close()

	a := advent.Advent{
		Input: file,
	}

	v, err := a.Day9Part1()
	if err != nil {
		t.Fatal(err)
	}

	if v != 13 {
		t.Errorf("expected 13, got %v", v)
	}
}

func TestDay9Part2(t *testing.T) {
	tests := map[string]int{"a": 1, "b": 36}
	for k, v := range tests {

		testFile := "data/day9" + k
		file, err := os.Open(testFile)
		if err != nil {
			t.Fatalf("could not open test data file %v", testFile)
		}
		defer file.Close()

		a := advent.Advent{
			Input: file,
		}

		got, err := a.Day9Part2()
		if err != nil {
			t.Fatal(err)
		}

		if got != v {
			t.Errorf("expected %v, got %v", v, got)
		}
	}
}
