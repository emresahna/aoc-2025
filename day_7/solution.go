package main

import (
	"fmt"
	"log"
	"strings"

	helper "github.com/emresahna/aoc-2025"
)

const START_POINT = "S"
const BEAM_POINT = "|"
const EMPTY_POINT = "."
const SPLITTER_POINT = "^"

func main() {
	input, err := helper.LoadFromFile("./day_7/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	password := 0
	firstArr := strings.Split(input[0], "")

	for i := 0; i < len(input)-1; i++ {
		secondArr := strings.Split(input[i+1], "")

		for j := 0; j < len(firstArr); j++ {
			if firstArr[j] == START_POINT {
				if secondArr[j] == EMPTY_POINT {
					secondArr[j] = BEAM_POINT
				} else {
					log.Fatal("unsupported input")
				}
			}

			if firstArr[j] == BEAM_POINT {
				if secondArr[j] == EMPTY_POINT {
					secondArr[j] = BEAM_POINT
				} else if secondArr[j] == SPLITTER_POINT {
					if j+1 != len(firstArr) {
						secondArr[j+1] = BEAM_POINT
					}

					if j-1 >= 0 {
						secondArr[j-1] = BEAM_POINT
					}
					password++
				} else if secondArr[j] != BEAM_POINT {
					log.Fatal("unsupported input")
				}
			}
		}

		firstArr = secondArr
	}

	fmt.Println("DAY 7 & PART 1 RESULT")
	fmt.Println(password)

	fmt.Println("============================")

	fmt.Println("DAY 7 & PART 2 RESULT")
	fmt.Println(password)
}
