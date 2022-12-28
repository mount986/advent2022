package day5

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

func ParseStacks(input string) [][]byte {
	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var stacks [][]byte

	for i, j := 0, 0; i < len(lines[len(lines)-1]); i, j = i+4, j+1 {
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

func ParseMove(move string) (int, int, int, error) {
	re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

	subs := re.FindStringSubmatch(move)

	count, err := strconv.Atoi(subs[1])
	if err != nil {
		return 0, 0, 0, err
	}

	from, err := strconv.Atoi(subs[2])
	if err != nil {
		return 0, 0, 0, err
	}

	to, err := strconv.Atoi(subs[3])
	if err != nil {
		return 0, 0, 0, err
	}

	return from - 1, to - 1, count, nil
}