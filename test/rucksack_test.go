package advent

import (
	"testing"
	"os"
	"github.com/mount986/advent2022/advent"
)

func TestRucksackPriorities(t *testing.T) {
	testFile := "data/rucksack"
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("could not open test data file %v", testFile )
	}
	defer file.Close()

	v, err := advent.RucksackPriorities(file)
	if err != nil {
		t.Fatal(err)
	}

	if v != 157 {
		t.Errorf("expected 157, got %v", v)
	}
}

func TestRucksackGroupPriorities(t *testing.T) {
	testFile := "data/rucksack"
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("could not open test data file %v", testFile )
	}
	defer file.Close()

	v, err := advent.RucksackGroupPriorities(file)
	if err != nil {
		t.Fatal(err)
	}

	if v != 70 {
		t.Errorf("expected 70, got %v", v)
	}
}