package utils

import (
	"sort"
	"testing"
)

//Benchmark for bubble sort
func BenchmarkBubbleSort(b *testing.B) {
	r := []int{1, 2, 3, 4}
	for i := 0; i <= b.N; i++ {
		BubbleSort(r)
	}
}

//Benchmark for go sort in built function
func BenchmarkSort(b *testing.B) {
	r := []int{1, 2, 3, 4}
	for i := 0; i <= b.N; i++ {
		sort.Ints(r)
	}
}
