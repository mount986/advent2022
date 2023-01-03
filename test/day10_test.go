package advent

import (
	"os"
	"testing"

	"github.com/mount986/advent2022/advent"
)

func TestDay10Part1(t *testing.T) {
	testFile := "data/day10"
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("could not open test data file %v", testFile)
	}
	defer file.Close()

	a := advent.Advent{
		Input: file,
	}

	v, err := a.Day10Part1()
	if err != nil {
		t.Fatal(err)
	}

	if v != 13140 {
		t.Errorf("expected 13140, got %v", v)
	}
}

func TestDay10Part2(t *testing.T) {
	testFile := "data/day10"
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("could not open test data file %v", testFile)
	}
	defer file.Close()

	a := advent.Advent{
		Input: file,
	}

	v, err := a.Day10Part2()
	if err != nil {
		t.Fatal(err)
	}

	if v != `
##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....` {
		t.Errorf("expected pattern does not match, got \n%v", v)
	}
}
