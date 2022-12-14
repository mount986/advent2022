package advent

import (
	"testing"
	"os"
	"github.com/mount986/advent2022/advent"
)

func TestRoshamboScoreSimple(t *testing.T) {
	testFile := "data/roshambo"
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("could not open test data file %v", testFile )
	}
	defer file.Close()

	score, err := advent.RoshamboScoreSimple(file)
	if err != nil {
		t.Fatal(err)
	}

	if score != 15 {
		t.Errorf("expected 15, got %v", score)
	}
}

func TestRoshamboScoreStrategy(t *testing.T) {
	testFile := "data/roshambo"
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("could not open test data file %v", testFile )
	}
	defer file.Close()

	score, err := advent.RoshamboScoreStrategy(file)
	if err != nil {
		t.Fatal(err)
	}

	if score != 12 {
		t.Errorf("expected 12, got %v", score)
	}	
}