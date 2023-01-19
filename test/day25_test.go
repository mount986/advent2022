package advent

import (
	"os"
	"testing"

	"github.com/mount986/advent2022/advent"
	"github.com/mount986/advent2022/advent/day25"
)

func TestDay25Part1(t *testing.T) {
	testFile := "data/day25"
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("could not open test data file %v", testFile)
	}
	defer file.Close()

	a := advent.Advent{
		Input: file,
	}

	v, err := a.Day25Part1()
	if err != nil {
		t.Fatal(err)
	}

	if v != "2=-1=0" {
		t.Errorf("expected 2=-1=0, got %v", v)
	}
}


func TestToDec(t *testing.T) {
	tests := map[string]int{
		"1=-0-2":        1747,
		"12111":         906,
		"2=0=":          198,
		"21":            11,
		"2=01":          201,
		"111":           31,
		"20012":         1257,
		"112":           32,
		"1=-1=":         353,
		"1-12":          107,
		"122":           37,
		"1":             1,
		"2":             2,
		"1=":            3,
		"1-":            4,
		"10":            5,
		"11":            6,
		"12":            7,
		"2=":            8,
		"2-":            9,
		"20":            10,
		"1=0":           15,
		"1-0":           20,
		"1=11-2":        2022,
		"1-0---0":       12345,
		"1121-1110-1=0": 314159265,
	}

	for k, v := range tests {
		got := day25.ToDec(k)
		
		if got != v {
			t.Errorf("Expected %v, got %v", v, got)
		}
	}
}

func TestToSnafu(t *testing.T) {
	tests := map[string]int{
		"1=-0-2":        1747,
		// "12111":         906,
		"2=0=":          198,
		// "21":            11,
		// "2=01":          201,
		// "111":           31,
		// "20012":         1257,
		// "112":           32,
		// "1=-1=":         353,
		// "1-12":          107,
		// "122":           37,
		// "1":             1,
		// "2":             2,
		// "1=":            3,
		// "1-":            4,
		// "10":            5,
		// "11":            6,
		// "12":            7,
		// "2=":            8,
		// "2-":            9,
		// "20":            10,
		// "1=0":           15,
		// "1-0":           20,
		// "1=11-2":        2022,
		// "1-0---0":       12345,
		"1121-1110-1=0": 314159265,
	}

	for k, v := range tests {
		got := day25.ToSnafu(v)
		
		if got != k {
			t.Errorf("Dec: %v; Expected %v, got %v", v, k, got)
		}
	}
}