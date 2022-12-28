package advent

import (
	"bufio"
	"bytes"
	"io"
	"strings"

	"github.com/mount986/advent2022/advent/day5"
	"golang.org/x/exp/slices"
)

func (a *Advent) Day5Part1() (string, error) {
	buf := bytes.NewBuffer(nil)
	
	io.Copy(buf, a.Input)
	input := buf.String()

	parts := strings.Split(input, "\n\n")

	stacks := day5.ParseStacks(parts[0])

	//loop through moves
	scanner := bufio.NewScanner(strings.NewReader(parts[1]))
	for scanner.Scan() {
		move := scanner.Text()
		from, to, count, err := day5.ParseMove(move)
		if err != nil {
			return "", err
		}

		for i := 0; i < count; i++ {
			box := stacks[from][i]
			stacks[to] = slices.Insert(stacks[to], 0 , box)
		}
	
		stacks[from] = slices.Delete(stacks[from], 0, count)
	}

	//gather top of stacks
	var results []byte
	for _, stack := range stacks {
		results = append(results, stack[0])
	}

	return string(results), nil
}

func (a *Advent) Day5Part2() (string, error) {

		buf := bytes.NewBuffer(nil)
	
	io.Copy(buf, a.Input)
	input := buf.String()

	parts := strings.Split(input, "\n\n")

	stacks := day5.ParseStacks(parts[0])

	//loop through moves
	scanner := bufio.NewScanner(strings.NewReader(parts[1]))
	for scanner.Scan() {
		move := scanner.Text()
		from, to, count, err := day5.ParseMove(move)
		if err != nil {
			return "", err
		}

		for i := 0; i < count; i++ {
			box := stacks[from][count -1 - i]
			stacks[to] = slices.Insert(stacks[to], 0 , box)
		}
	
		stacks[from] = slices.Delete(stacks[from], 0, count)
	}

	//gather top of stacks
	var results []byte
	for _, stack := range stacks {
		results = append(results, stack[0])
	}

	return string(results), nil
}
