package main

import (
	"fmt"
	"log"
	"strconv"

	helper "github.com/emresahna/aoc-2025"
)

const INITIAL_NUMBER = 50

func main() {
	input, err := helper.LoadInput("https://adventofcode.com/2025/day/1/input")
	if err != nil {
		log.Fatal(err)
	}

	password := 0
	curr := INITIAL_NUMBER

	for _, element := range input {
		numb, err := strconv.Atoi(element[1:])
		if err != nil {
			log.Fatal("error while parsing elements")
		}

		switch string(element[0]) {
		case "L":
			curr -= numb
		case "R":
			curr += numb
		default:
			log.Fatal("unknown direction")
		}

		curr = curr % 100

		if curr == 0 {
			password++
		}
	}

	fmt.Println("DAY 1 & PART 1 RESULT")
	fmt.Println(password)

	fmt.Println("============================")

	fmt.Println("DAY 1 & PART 2 RESULT")
	fmt.Println(password)
}
