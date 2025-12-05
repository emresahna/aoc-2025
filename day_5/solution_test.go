package main

import (
	"log"
	"testing"

	helper "github.com/emresahna/aoc-2025"
)

const inputUrl = "https://adventofcode.com/2025/day/5/input"

func BenchmarkUnoptimizedFunc(b *testing.B) {
	input, err := helper.LoadInput(inputUrl)
	if err != nil {
		log.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = NotOptimizedSolution(input)
	}
}

func BenchmarkOptimizedFunc(b *testing.B) {
	input, err := helper.LoadInput(inputUrl)
	if err != nil {
		log.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = OptimizedSolution(input)
	}
}
