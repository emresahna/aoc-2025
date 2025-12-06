package main

import (
	"fmt"
	"log"
	"strconv"

	helper "github.com/emresahna/aoc-2025"
)

const DIGIT_PER_JOLTAGE = 12

func main() {
	input, err := helper.LoadFromSource("https://adventofcode.com/2025/day/3/input")
	if err != nil {
		log.Fatal(err)
	}

	password := 0

	for _, battery := range input {
		left := 0
		leftIndex := 0
		for index, bank := range battery[:len(battery)-1] {
			nmb, err := strconv.Atoi(string(bank))
			if err != nil {
				log.Fatal(err)
			}

			if nmb > left {
				left = nmb
				leftIndex = index
			}
		}

		right := 0
		for index := leftIndex + 1; index < len(battery); index++ {
			nmb, err := strconv.Atoi(string(battery[index]))
			if err != nil {
				log.Fatal(err)
			}

			if nmb > right {
				right = nmb
			}
		}

		password += right + (left * 10)
	}

	fmt.Println("DAY 3 & PART 1 RESULT")
	fmt.Println(password)

	fmt.Println("============================")

	fmt.Println("DAY 3 & PART 2 RESULT")
	fmt.Println(password)
}
