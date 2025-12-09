package main

import (
	"fmt"
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

	fmt.Println(actual)
	if actual != expected {
		t.Error("example test case failed")
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
		t.Error("example test case failed")
	}
}
