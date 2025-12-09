package main

import (
	"cmp"
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"

	helper "github.com/emresahna/aoc-2025"
)

func main() {
	input, err := helper.LoadFromFile("./day_8/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	password, err := JunctionMatcher(input, 1000, 3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DAY 8 & PART 1 RESULT")
	fmt.Println(password)

	fmt.Println("============================")

	password, err = JunctionMatcher(input, 0, 0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DAY 8 & PART 2 RESULT")
	fmt.Println(password)
}

type Pair struct {
	P1       string
	P2       string
	Distance float64
}

func JunctionMatcher(input []string, pairCount int, highestCount int) (int, error) {
	var allPairs []map[string]struct{}
	for _, i := range input {
		allPairs = append(allPairs, map[string]struct{}{i: {}})
	}

	var anyArr []Pair
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			anyArr = append(anyArr, Pair{
				P1:       input[i],
				P2:       input[j],
				Distance: calculateOclide(input[i], input[j]),
			})
		}
	}

	slices.SortFunc(anyArr, func(a, b Pair) int {
		return cmp.Compare(a.Distance, b.Distance)
	})

	stopLatestPair := false
	if pairCount == 0 && highestCount == 0 {
		pairCount = len(anyArr)
		stopLatestPair = true
	}

	for i := 0; i < pairCount; i++ {
		p1 := anyArr[i].P1
		p2 := anyArr[i].P2

		p1Found := false
		p2Found := false
		p1FoundIndex := 0
		p2FoundIndex := 0

		for index, node := range allPairs {
			if p1Found && p2Found {
				break
			}

			_, foundP1 := node[p1]
			_, foundP2 := node[p2]

			if foundP1 {
				p1Found = true
				p1FoundIndex = index
			}

			if foundP2 {
				p2Found = true
				p2FoundIndex = index
			}
		}

		if p1FoundIndex != p2FoundIndex {
			for key := range allPairs[p1FoundIndex] {
				allPairs[p2FoundIndex][key] = struct{}{}
			}
			allPairs = append(allPairs[:p1FoundIndex], allPairs[p1FoundIndex+1:]...)

			if stopLatestPair && len(allPairs) == 1 {
				p1Coordinates := strings.Split(anyArr[i].P1, ",")
				p2Coordinates := strings.Split(anyArr[i].P2, ",")

				p1X, err := strconv.Atoi(p1Coordinates[0])
				if err != nil {
					log.Fatal(err)
				}

				p2X, err := strconv.Atoi(p2Coordinates[0])
				if err != nil {
					log.Fatal(err)
				}

				return p1X * p2X, nil
			}
		}
	}

	slices.SortFunc(allPairs, func(a, b map[string]struct{}) int {
		return cmp.Compare(len(b), len(a))
	})

	result := 1
	for i := range highestCount {
		result *= len(allPairs[i])
	}

	return result, nil
}

func calculateOclide(p1 string, p2 string) float64 {
	p1Coordinates := strings.Split(p1, ",")
	p2Coordinates := strings.Split(p2, ",")

	p1X, err := strconv.Atoi(p1Coordinates[0])
	if err != nil {
		log.Fatal(err)
	}

	p1Y, err := strconv.Atoi(p1Coordinates[1])
	if err != nil {
		log.Fatal(err)
	}

	p1Z, err := strconv.Atoi(p1Coordinates[2])
	if err != nil {
		log.Fatal(err)
	}

	p2X, err := strconv.Atoi(p2Coordinates[0])
	if err != nil {
		log.Fatal(err)
	}

	p2Y, err := strconv.Atoi(p2Coordinates[1])
	if err != nil {
		log.Fatal(err)
	}

	p2Z, err := strconv.Atoi(p2Coordinates[2])
	if err != nil {
		log.Fatal(err)
	}

	xTotalDiff := math.Pow(float64(p2X-p1X), 2)
	yTotalDiff := math.Pow(float64(p2Y-p1Y), 2)
	zTotalDiff := math.Pow(float64(p2Z-p1Z), 2)

	totalDiff := math.Abs(xTotalDiff + yTotalDiff + zTotalDiff)

	distance := math.Pow(totalDiff, 0.5)

	return distance
}
