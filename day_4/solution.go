package main

import (
	"fmt"
	"log"

	helper "github.com/emresahna/aoc-2025"
)

const ACCESSABLE_ROLL_BOUNDY = 4
const ROLL = '@'
const EMPTY = '.'

func main() {
	input, err := helper.LoadInput("https://adventofcode.com/2025/day/4/input")
	if err != nil {
		log.Fatal(err)
	}

	password := 0

	for row := 0; row < len(input); row++ {
		for tile := 0; tile < len(input[row]); tile++ {
			if input[row][tile] == EMPTY {
				continue
			}

			c := 0

			// up
			if row > 0 && input[row-1][tile] == ROLL {
				c++
			}

			// down
			if row < len(input)-1 && input[row+1][tile] == ROLL {
				c++
			}

			// left
			if tile > 0 && input[row][tile-1] == ROLL {
				c++
			}

			// right
			if tile < len(input[row])-1 && input[row][tile+1] == ROLL {
				c++
			}

			// up left
			if row > 0 && tile > 0 && input[row-1][tile-1] == ROLL {
				c++
			}

			// up right
			if row > 0 && tile < len(input[row])-1 && input[row-1][tile+1] == ROLL {
				c++
			}

			// down left
			if row < len(input)-1 && tile > 0 && input[row+1][tile-1] == ROLL {
				c++
			}

			// down right
			if row < len(input)-1 && tile < len(input[row])-1 && input[row+1][tile+1] == ROLL {
				c++
			}

			if c < ACCESSABLE_ROLL_BOUNDY {
				password++
			}
		}
	}

	fmt.Println("DAY 4 & PART 1 RESULT")
	fmt.Println(password)

	fmt.Println("============================")

	fmt.Println("DAY 4 & PART 2 RESULT")
	fmt.Println(password)
}
