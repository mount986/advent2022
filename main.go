package main

import (
	"fmt"
	"os"

	"github.com/mount986/advent2022/advent"
)

func main() {
	fmt.Println("Running!")

	if len(os.Args[1:]) < 1 {
		all()
		return
	}
	f := os.Args[1]
	switch f {
	case "1", "calories":
		day1()
	case "2", "roshambo":
		day2()
	case "3", "rucksack":
		day3()
	default:
		panic(fmt.Errorf("invalid function"))
	}
}

func day1() {
	fmt.Println("Running Calories!")
	file, err := os.Open("./data/calories")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	max, err := advent.CaloriesMax(file)
	if err != nil {
		panic(err)
	}

	file.Seek(0, 0)

	max3, err := advent.CaloriesMax3(file)
	if err != nil {
		panic(err)
	}

	fmt.Println("calories max: ", max)
	fmt.Println("calories max3: ", max3)
}

func day2() {
	file, err := os.Open("./data/roshambo")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	score1, err := advent.RoshamboScoreSimple(file)
	if err != nil {
		panic(err)
	}

	file.Seek(0, 0)

	score2, err := advent.RoshamboScoreStrategy(file)
	if err != nil {
		panic(err)
	}

	fmt.Println("roshambo simple score: ", score1)
	fmt.Println("roshambo strategic score: ", score2)
}

func day3() {
	file, err := os.Open("./data/rucksack")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	score1, err := advent.RucksackPriorities(file)
	if err != nil {
		panic(err)
	}

	file.Seek(0, 0)

	score2, err := advent.RucksackGroupPriorities(file)
	if err != nil {
		panic(err)
	}

	fmt.Println("rucksack priorities: ", score1)
	fmt.Println("roshambo group priorities: ", score2)
}

func all() {
	day1()
	day2()
	day3()
}
