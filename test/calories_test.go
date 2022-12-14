package advent

import (
	"os"
	"testing"
	"github.com/mount986/advent2022/advent"
)

func TestCaloriesMax(t *testing.T) {
	testFile := "data/calories"
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("could not open test data file %v", testFile )
	}
	defer file.Close()

	max, err := advent.CaloriesMax(file)
	if err != nil {
		t.Fatal(err)
	}

	if max != 24000 {
		t.Errorf("expected 2400, got %v", max)
	}
}

func TestCaloriesMax3(t *testing.T) {
	testFile := "data/calories"
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("could not open test data file %v", testFile )
	}
	defer file.Close()

	max, err := advent.CaloriesMax3(file)
	if err != nil {
		t.Fatal(err)
	}

	if max != 45000 {
		t.Errorf("expected 45000, got %v", max)
	}
}