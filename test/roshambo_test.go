package advent

import (
	"os"
	"testing"

	"github.com/mount986/advent2022/advent"
)

func TestDay2Part1(t *testing.T) {
	testFile := "data/roshambo"
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("could not open test data file %v", testFile)
	}
	defer file.Close()

	score, err := advent.Day2Part1(file)
	if err != nil {
		t.Fatal(err)
	}

	if score != 15 {
		t.Errorf("expected 15, got %v", score)
	}
}

func TestDay2Part2(t *testing.T) {
	testFile := "data/roshambo"
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("could not open test data file %v", testFile)
	}
	defer file.Close()

	score, err := advent.Day2Part2(file)
	if err != nil {
		t.Fatal(err)
	}

	if score != 12 {
		t.Errorf("expected 12, got %v", score)
	}
}
