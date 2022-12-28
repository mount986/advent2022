package advent

import (
	"bufio"
	"bytes"
	"io"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func (a *Advent) Day5Part1() (string, error) {
	buf := bytes.NewBuffer(nil)
	
	io.Copy(buf, a.Input)
	input := buf.String()

	parts := strings.Split(input, "\n\n")

	stacks := parseStacks(parts[0])

	//loop through moves
	scanner := bufio.NewScanner(strings.NewReader(parts[1]))
	for scanner.Scan() {
		move := scanner.Text()
		from, to, count, err := parseMove(move)
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

	stacks := parseStacks(parts[0])

	//loop through moves
	scanner := bufio.NewScanner(strings.NewReader(parts[1]))
	for scanner.Scan() {
		move := scanner.Text()
		from, to, count, err := parseMove(move)
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

func parseStacks(input string) ([][]byte) {
	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())		
	}

	var stacks [][]byte

	for i, j := 0, 0; i < len(lines[len(lines) - 1]); i, j = i + 4, j + 1{
		var stack []byte
		for _, line := range lines {
			if line[i] == '[' {
				stack = append(stack, line[i+1])
			}
		}
		stacks = append(stacks, stack)
	}

	return stacks
}

func parseMove(move string) (int, int, int, error) {
	re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

	subs := re.FindStringSubmatch(move)

	count, err := strconv.Atoi(subs[1])
	if err != nil {
		return 0,0,0,err
	}

	from, err := strconv.Atoi(subs[2])
	if err != nil {
		return 0,0,0,err
	}

	to, err := strconv.Atoi(subs[3])
	if err != nil {
		return 0,0,0,err
	}

	return from - 1, to - 1, count, nil
}
