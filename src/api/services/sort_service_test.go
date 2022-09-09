package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getElements(n int) []int {
	result := make([]int, n)
	j := 0
	for i := n - 1; i > 0; i-- {
		result[i] = i
		j++
	}
	return result
}
func TestSort(t *testing.T) {
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	Sort(elements)

	// Validation
	assert.EqualValues(t, 9, elements[0], "first element should be 9")
	assert.EqualValues(t, 0, elements[len(elements)-1], "last element should be 0")
}

func TestSortMoreThan10000(t *testing.T) {
	elements := getElements(10001)
	Sort(elements)

	// Validation
	assert.EqualValues(t, 0, elements[0], "first element should be 0")
	assert.EqualValues(t, 10000, elements[len(elements)-1], "last element should be 10000")
}
