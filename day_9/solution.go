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

type Point struct {
	X int
	Y int
}

type Rectangel struct {
	Corner1 Point
	Corner2 Point
	Area    int
}

func main() {
	input, err := helper.LoadFromFile("./day_9/example_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	password, err := FindLargestRectangleArea(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DAY 9 & PART 1 RESULT")
	fmt.Println(password)

	fmt.Println("============================")

	fmt.Println("DAY 9 & PART 2 RESULT")
	fmt.Println(password)
}

func FindLargestRectangleArea(input []string) (int, error) {
	points, err := convertInputToPoints(input)
	if err != nil {
		return 0, err
	}

	possibleRectangles, err := createPossibleRectangles(points)
	if err != nil {
		return 0, err
	}

	slices.SortFunc(possibleRectangles, func(x Rectangel, y Rectangel) int {
		return cmp.Compare(y.Area, x.Area)
	})

	return possibleRectangles[0].Area, nil
}

func createPossibleRectangles(points []Point) ([]Rectangel, error) {
	pointsLen := len(points)
	possibleRectangleLen := (pointsLen * (pointsLen - 1)) / 2
	rectangels := make([]Rectangel, possibleRectangleLen)
	placed := 0

	for iIndex, i := range points {
		for j := iIndex + 1; j < len(points); j++ {
			area, err := calculateArea(i, points[j])
			if err != nil {
				return nil, err
			}

			rectangels[placed] = Rectangel{
				Corner1: i,
				Corner2: points[j],
				Area:    area,
			}
			placed++
		}
	}

	return rectangels, nil
}

func calculateArea(corner1 Point, corner2 Point) (int, error) {
	diffX := math.Abs(float64(corner2.X-corner1.X)) + 1
	diffY := math.Abs(float64(corner2.Y-corner1.Y)) + 1

	return int(diffX * diffY), nil
}

func convertInputToPoints(input []string) ([]Point, error) {
	points := make([]Point, len(input))
	for iIndex, i := range input {
		point1X, point1Y, err := parseCoordToInt(i)
		if err != nil {
			return nil, err
		}

		points[iIndex] = Point{X: point1X, Y: point1Y}
	}

	return points, nil
}

func parseCoordToInt(corner string) (int, int, error) {
	cornerCoord := strings.Split(corner, ",")

	cornerX, err := strconv.Atoi(cornerCoord[0])
	if err != nil {
		return 0, 0, nil
	}

	cornerY, err := strconv.Atoi(cornerCoord[1])
	if err != nil {
		return 0, 0, nil
	}

	return cornerX, cornerY, nil
}
