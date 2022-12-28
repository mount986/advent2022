package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/mount986/advent2022/advent"
)

func main() {
	fmt.Println("Running!")

	if len(os.Args[1:]) < 1 {
		all()
		return
	}
	run(os.Args[1])

}

func run(day string) {
	fmt.Println("Running Day ", day)

	file, err := os.Open("./data/day" + day)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	a := advent.Advent{Input: file}

	results := reflect.ValueOf(&a).MethodByName("Day" + day + "Part1").Call([]reflect.Value{})
	if !results[1].IsNil() {
		panic(err)
	}

	fmt.Printf("Day %v, Part 1: %v\n", day, results[0])

	file.Seek(0, 0)

	results = reflect.ValueOf(&a).MethodByName("Day" + day + "Part2").Call([]reflect.Value{})
	if !results[1].IsNil() {
		panic(err)
	}

	fmt.Printf("Day %v, Part 2: %v\n", day, results[0])
}

func all() {
	for x :=1; x <= 25; x++ {
		run(fmt.Sprint(x))
	}
}
