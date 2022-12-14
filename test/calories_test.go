package advent

import (
	"os"
	"testing"

	"github.com/mount986/advent2022/advent"
)

func TestDay1Part1(t *testing.T) {
	testFile := "data/calories"
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("could not open test data file %v", testFile)
	}
	defer file.Close()

	max, err := advent.Day1Part1(file)
	if err != nil {
		t.Fatal(err)
	}

	if max != 24000 {
		t.Errorf("expected 2400, got %v", max)
	}
}

func TestDay1Part2(t *testing.T) {
	testFile := "data/calories"
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("could not open test data file %v", testFile)
	}
	defer file.Close()

	max, err := advent.Day1Part2(file)
	if err != nil {
		t.Fatal(err)
	}

	if max != 45000 {
		t.Errorf("expected 45000, got %v", max)
	}
}
