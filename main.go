package main

import (
	"fmt"
	"os"

	"github.com/mount986/advent2022/calories"
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

		max, err := calories.CaloriesMax(file)
		if err != nil {
			panic(err)
		}

		file.Seek(0, 0)

		max3, err := calories.CaloriesMax3(file)
		if err != nil {
			panic(err)
		}

		fmt.Println("calories max: ", max)
		fmt.Println("calories max3: ", max3)
}

func all() {
	day1()
}