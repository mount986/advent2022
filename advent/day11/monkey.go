package day11

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

type Monkey struct {
	Items     []int
	Operation func(int) int
	Test      int
	TMonkey   int
	FMonkey   int
	Count     int
}

func NewMonkey(scanner *bufio.Scanner) (Monkey, error) {
	var m Monkey

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		if line[:17] == "  Starting items:" {
			items := strings.Split(line[18:], ", ")
			for _, item := range items {
				i, err := strconv.Atoi(item)
				if err != nil {
					return m, err
				}
				m.Items = append(m.Items, i)
			}

		} else if line[:12] == "  Operation:" {
			re, err := regexp.Compile(`new = old ([\+\*]) (\w+)`)
			if err != nil {
				return m, err
			}
			matches := re.FindStringSubmatch(line[13:])

			switch matches[1] {
			case "*":
				if matches[2] == "old" {
					m.Operation = func(old int) int { return old * old }
				} else {
					multiplicand, err := strconv.Atoi(matches[2])
					if err != nil {
						return m, err
					}
					m.Operation = func(old int) int { return old * multiplicand }
				}
			case "+":
				if matches[2] == "old" {
					m.Operation = func(old int) int { return old + old }
				} else {
					addend, err := strconv.Atoi(matches[2])
					if err != nil {
						return m, err
					}
					m.Operation = func(old int) int { return old + addend }
				}
			}

		} else if line[:7] == "  Test:" {
			re, err := regexp.Compile(`divisible by (\d+)`)
			if err != nil {
				return m, err
			}
			matches := re.FindStringSubmatch(line[8:])
			Test, err := strconv.Atoi(matches[1])
			if err != nil {
				return m, err
			}
			m.Test = Test

		} else if line[:12] == "    If true:" {
			re, err := regexp.Compile(`throw to monkey (\d+)`)
			if err != nil {
				return m, err
			}
			matches := re.FindStringSubmatch(line[13:])
			v, err := strconv.Atoi(matches[1])
			if err != nil {
				return m, err
			}
			m.TMonkey = v

		} else if line[:13] == "    If false:" {
			re, err := regexp.Compile(`throw to monkey (\d+)`)
			if err != nil {
				return m, err
			}
			matches := re.FindStringSubmatch(line[14:])
			v, err := strconv.Atoi(matches[1])
			if err != nil {
				return m, err
			}
			m.FMonkey = v

		}
	}

	m.Count = 0

	return m, nil
}

func (m *Monkey) Investigate(monkeys []Monkey, relief bool, lcm int) {
	for index := range m.Items {
		m.Count++
		m.Items[index] = m.Operation(m.Items[index])
		if relief {
			m.Items[index] /= 3
		}
		m.Items[index] = m.Items[index] % lcm

		if m.Items[index] % m.Test == 0 {
			monkeys[m.TMonkey].Items = append(monkeys[m.TMonkey].Items, m.Items[index])
		} else {
			monkeys[m.FMonkey].Items = append(monkeys[m.FMonkey].Items, m.Items[index])
		}
	}

	m.Items = make([]int, 0)
}
