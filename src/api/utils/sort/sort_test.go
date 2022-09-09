package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	// init
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	fmt.Println(elements)
	// execution
	BubbleSort(elements)

	// Validation
	if elements[0] != 9 {
		t.Error("first element should be 9")
	}
	if elements[len(elements)-1] != 0 {
		t.Error("last element should be 0")
	}
	fmt.Println(elements)
}

func TestSort(t *testing.T) {
	// init
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	fmt.Println(elements)
	// execution
	Sort(elements)

	// Validation
	if elements[0] != 0 {
		t.Error("first element should be 0")
	}
	if elements[len(elements)-1] != 9 {
		t.Error("last element should be 9")
	}
	fmt.Println(elements)
}

func getElements(n int) []int {
	result := make([]int, n)
	j := 0
	for i := n - 1; i > 0; i-- {
		result[i] = i
		j++
	}
	return result
}

func BenchmarkBubbleSort(b *testing.B) {
	// elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	elements := getElements(10000)
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkSort(b *testing.B) {
	// elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	elements := getElements(10000)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
