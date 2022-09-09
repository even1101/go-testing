package services

import "testing"

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
	if elements[0] != 9 {
		t.Error("first element should be 9")
	}
	if elements[len(elements)-1] != 0 {
		t.Error("last element should be 0")
	}
}

func TestSortMoreThan10000(t *testing.T) {
	elements := getElements(10001)
	Sort(elements)

	// Validation
	if elements[0] != 10000 {
		t.Error("first element should be 10000")
	}
	if elements[len(elements)-1] != 0 {
		t.Error("last element should be 0")
	}
}
