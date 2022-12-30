package advent

import (
	"os"
	"testing"

	"github.com/mount986/advent2022/advent"
)

func TestDay6Part1(t *testing.T) {
	tests := map[string]int{"a": 7, "b": 5, "c": 6, "d": 10, "e": 11}
	for k, v := range tests {

		testFile := "data/day06" + k
		file, err := os.Open(testFile)
		if err != nil {
			t.Fatalf("could not open test data file %v", testFile)
		}
		defer file.Close()

		a := advent.Advent{
			Input: file,
		}

		got, err := a.Day6Part1()
		if err != nil {
			t.Fatal(err)
		}

		if got != v {
			t.Errorf("expected %v, got %v", v, got)
		}
	}
}

func TestDay6Part2(t *testing.T) {
	tests := map[string]int{"a": 19, "b": 23, "c": 23, "d": 29, "e": 26}
	for k, v := range tests {

		testFile := "data/day06" + k
		file, err := os.Open(testFile)
		if err != nil {
			t.Fatalf("could not open test data file %v", testFile)
		}
		defer file.Close()

		a := advent.Advent{
			Input: file,
		}

		got, err := a.Day6Part2()
		if err != nil {
			t.Fatal(err)
		}

		if got != v {
			t.Errorf("expected %v, got %v", v, got)
		}
	}
}
