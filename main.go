package main

import (
	"fmt"
	"go-testing/utils/sort"
)

func main() {
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	fmt.Println(elements)
	// execution
	sort.BubbleSort(elements)
	fmt.Println(elements)

}
