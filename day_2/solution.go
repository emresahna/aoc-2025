package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	helper "github.com/emresahna/aoc-2025"
)

func main() {
	input, err := helper.LoadInput("https://adventofcode.com/2025/day/2/input")
	if err != nil {
		log.Fatal(err)
	}

	password := 0

	for pair := range strings.SplitSeq(input[0], ",") {
		nmbs := strings.Split(pair, "-")

		nmb0, err := strconv.Atoi(nmbs[0])
		if err != nil {
			log.Fatal(err)
		}

		nmb1, err := strconv.Atoi(nmbs[1])
		if err != nil {
			log.Fatal(err)
		}

		for index := nmb0; index < nmb1; index++ {
			s := strconv.Itoa(index)
			l := len(s)
			left := 0
			right := l / 2

			if l%2 != 0 {
				continue
			}

			valid := true
			for i := range right {
				if s[left+i] != s[right+i] {
					valid = false
					break
				}
			}

			if valid {
				password += index
			}
		}

	}

	fmt.Println("DAY 2 & PART 1 RESULT")
	fmt.Println(password)

	fmt.Println("============================")

	fmt.Println("DAY 2 & PART 2 RESULT")
	fmt.Println(password)
}
