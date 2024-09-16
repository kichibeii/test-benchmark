package main

import (
	"testing"

	"github.com/kichibeii/benchmark-tool/benchmarktool"
	"golang.org/x/exp/rand"
)

func makeRandomNumberSlice() []int {
	LENGTH := 1000 // Set the desired length of the random slice
	numbers := make([]int, LENGTH)
	for i := 0; i < LENGTH; i++ {
		numbers[i] = rand.Intn(1000)
	}
	return numbers
}

func BenchmarkBubbleSort(b *testing.B) {
	params := []interface{}{makeRandomNumberSlice()}
	benchmarktool.FunctionBenchmark(b, "BubbleSort", BubbleSort, params)
}

func BenchmarkQuickSort(b *testing.B) {
	params := []interface{}{makeRandomNumberSlice()}
	benchmarktool.FunctionBenchmark(b, "QuickSort", QuickSort, params)
}
