package advent

import (
	"os"
	"testing"

	"github.com/mount986/advent2022/advent"
)

func TestDay3Part1(t *testing.T) {
	testFile := "data/rucksack"
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("could not open test data file %v", testFile)
	}
	defer file.Close()

	v, err := advent.Day3Part1(file)
	if err != nil {
		t.Fatal(err)
	}

	if v != 157 {
		t.Errorf("expected 157, got %v", v)
	}
}

func TestDay3Part2(t *testing.T) {
	testFile := "data/rucksack"
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("could not open test data file %v", testFile)
	}
	defer file.Close()

	v, err := advent.Day3Part2(file)
	if err != nil {
		t.Fatal(err)
	}

	if v != 70 {
		t.Errorf("expected 70, got %v", v)
	}
}
