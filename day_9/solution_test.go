package main

import (
	"log"
	"testing"

	helper "github.com/emresahna/aoc-2025"
)

func TestExampleCase_Part1(t *testing.T) {
	input, err := helper.LoadFromFile("example_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	expected := 50

	actual, err := FindLargestRectangleArea(input)
	if err != nil {
		log.Fatal(err)
	}

	if expected != actual {
		t.Errorf("example test case failed expected: %d actual: %d", expected, actual)
	}
}
