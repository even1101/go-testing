package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	// init
	elements := getElements(10)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 0, elements[0])
	assert.EqualValues(t, 9, elements[len(elements)-1])
	// execution
	BubbleSort(elements)
	// Validation

	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 9, elements[0])
	assert.EqualValues(t, 0, elements[len(elements)-1])
}

func TestSort(t *testing.T) {
	// init
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	// execution
	Sort(elements)
	// Validation
	assert.EqualValues(t, 0, elements[0], "first element should be 0")
	assert.EqualValues(t, 9, elements[len(elements)-1], "last element should be 9")
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
