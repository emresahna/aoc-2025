package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	helper "github.com/emresahna/aoc-2025"
)

const ID_FINISH_IDENTIFIER = ""

func main() {
	input, err := helper.LoadInput("https://adventofcode.com/2025/day/5/input")
	if err != nil {
		log.Fatal(err)
	}

	password := OptimizedSolution(input)

	fmt.Println("DAY 5 & PART 1 RESULT")
	fmt.Println(password)

	fmt.Println("============================")

	fmt.Println("DAY 5 & PART 2 RESULT")
	fmt.Println(password)
}

func OptimizedSolution(workingInput []string) int {
	input := make([]string, len(workingInput))
	copy(input, workingInput)

	i := 0
	for input[i] != ID_FINISH_IDENTIFIER {
		i++
	}

	id_ranges := input[:i]

	for index, j := range id_ranges {
		numb1s := strings.Split(j, "-")

		lowest1, err := strconv.Atoi(numb1s[0])
		if err != nil {
			log.Fatal(err)
		}

		higest1, err := strconv.Atoi(numb1s[1])
		if err != nil {
			log.Fatal(err)
		}

		for index2, k := range id_ranges {
			if k == "" {
				continue
			}

			numb2s := strings.Split(k, "-")

			lowest2, err := strconv.Atoi(numb2s[0])
			if err != nil {
				log.Fatal(err)
			}

			higest2, err := strconv.Atoi(numb2s[1])
			if err != nil {
				log.Fatal(err)
			}

			if (lowest2 < higest1 && lowest2 > lowest1) || (higest2 > lowest1 && higest2 < higest1) {
				lwst := lowest1
				if lowest2 < lowest1 {
					lwst = lowest2
				}

				hgst := higest1
				if higest2 > higest1 {
					hgst = higest2
				}

				id_ranges[index2] = fmt.Sprintf("%d-%d", lwst, hgst)
				id_ranges[index] = ""
			}
		}
	}

	var optmzedArr []string
	for _, optmzdItem := range id_ranges {
		if optmzdItem != "" {
			optmzedArr = append(optmzedArr, optmzdItem)
		}
	}

	password := 0
	ids := input[i+1:]

	for _, idStr := range ids {
		idInt, err := strconv.Atoi(idStr)
		if err != nil {
			log.Fatal(err)
		}

		for _, id_rangeStr := range optmzedArr {
			id_rangeInt := strings.Split(id_rangeStr, "-")

			lowest, err := strconv.Atoi(id_rangeInt[0])
			if err != nil {
				log.Fatal(err)
			}

			highest, err := strconv.Atoi(id_rangeInt[1])
			if err != nil {
				log.Fatal(err)
			}

			if idInt > lowest && idInt < highest {
				password++
				break
			}
		}
	}

	return password
}

func NotOptimizedSolution(workingInput []string) int {
	input := make([]string, len(workingInput))
	copy(input, workingInput)

	i := 0
	for input[i] != ID_FINISH_IDENTIFIER {
		i++
	}

	id_ranges := input[:i]
	ids := input[i+1:]

	password := 0

	for _, idStr := range ids {
		idInt, err := strconv.Atoi(idStr)
		if err != nil {
			log.Fatal(err)
		}

		for _, id_rangeStr := range id_ranges {
			id_rangeInt := strings.Split(id_rangeStr, "-")

			lowest, err := strconv.Atoi(id_rangeInt[0])
			if err != nil {
				log.Fatal(err)
			}

			highest, err := strconv.Atoi(id_rangeInt[1])
			if err != nil {
				log.Fatal(err)
			}

			if idInt > lowest && idInt < highest {
				password++
				break
			}
		}
	}

	return password
}
