package day11

import (
	"bufio"
	"sort"
)

var lcm int

func Init(scanner *bufio.Scanner) ([]Monkey, error) {
	var monkeys []Monkey
	for scanner.Scan() {
		monkey, err := NewMonkey(scanner)
		if err != nil {
			return monkeys, err
		}
		monkeys = append(monkeys, monkey)
	}

	lcm = LCM(monkeys)

	return monkeys, nil
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(monkeys []Monkey) int {
	result := 1

	for _, monkey := range monkeys {
		result = result * monkey.Test / GCD(result, monkey.Test)
	}

	return result
}

func Run(monkeys []Monkey, rounds int, relief bool) (int, error) {
	for round := 0; round < rounds; round++ {
		for index := range monkeys {
			monkeys[index].Investigate(monkeys, relief, lcm)
		}
	}

	sortf := func(i, j int) bool { return monkeys[i].Count > monkeys[j].Count }

	sort.Slice(monkeys, sortf)

	return monkeys[0].Count * monkeys[1].Count, nil
}
