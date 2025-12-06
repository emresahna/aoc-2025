package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	helper "github.com/emresahna/aoc-2025"
)

func main() {
	input, err := helper.LoadFromFile("./day_6/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lineArr := make([][]string, len(input))
	for index, line := range input {
		lineArr[index] = extractLine(line)
	}

	arrLen := len(lineArr)
	operatorBoundry := len(lineArr[arrLen-1])
	password := 0

	for i := range operatorBoundry {
		var processColumns []string
		for j := range arrLen - 1 {
			processColumns = append(processColumns, lineArr[j][i])
		}

		res, err := calculateLine(processColumns, lineArr[arrLen-1][i])
		if err != nil {
			log.Fatal(err)
		}

		password += res
	}

	fmt.Println("DAY 5 & PART 1 RESULT")
	fmt.Println(password)

	fmt.Println("============================")

	fmt.Println("DAY 5 & PART 2 RESULT")
	fmt.Println(password)
}

func calculateLine(numbersStr []string, operator string) (int, error) {
	if operator == "+" {
		res := 0
		for _, i := range numbersStr {
			nmb, err := strconv.Atoi(i)
			if err != nil {
				return 0, err
			}

			res += nmb
		}
		return res, nil
	}

	if operator == "*" {
		res := 1
		for _, i := range numbersStr {
			nmb, err := strconv.Atoi(i)
			if err != nil {
				return 0, err
			}

			res *= nmb
		}
		return res, nil
	}

	return 0, fmt.Errorf("unsupported operator")
}

func extractLine(line string) []string {
	var res []string
	var found bool
	var each strings.Builder

	index := 0
	for index != len(line) {

		if line[index] != ' ' {
			each.WriteByte(line[index])

			if !found {
				found = true
			}
		}

		if found && line[index] == ' ' {
			res = append(res, each.String())

			found = false
			each.Reset()
		}

		index++
	}

	if each.Len() != 0 {
		res = append(res, each.String())
	}

	return res
}
