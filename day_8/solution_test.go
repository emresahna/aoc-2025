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

	expected := 40

	actual, err := JunctionMatcher(input, 10, 3)
	if err != nil {
		log.Fatal(err)
	}

	if actual != expected {
		t.Errorf("example test case failed expected: %d actual: %d", expected, actual)
	}
}

func TestExampleCase_Part2(t *testing.T) {
	input, err := helper.LoadFromFile("example_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	expected := 25272

	actual, err := JunctionMatcher(input, 0, 0)
	if err != nil {
		log.Fatal(err)
	}

	if actual != expected {
		t.Errorf("example test case failed expected: %d actual: %d", expected, actual)
	}
}
