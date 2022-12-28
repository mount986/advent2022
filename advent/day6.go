package advent

import (
	"bytes"
	"io"

	"github.com/mount986/advent2022/advent/day6"
)

func (a *Advent) Day6Part1() (int, error) {
	buf := bytes.NewBuffer(nil)

	io.Copy(buf, a.Input)
	input := buf.Bytes()	

	return day6.FindUnique(input, 4)
}

func (a *Advent) Day6Part2() (int, error) {
	buf := bytes.NewBuffer(nil)

	io.Copy(buf, a.Input)
	input := buf.Bytes()	

	return day6.FindUnique(input, 14)
}