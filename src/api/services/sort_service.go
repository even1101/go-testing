package services

import "go-testing/src/api/utils/sort"

func Sort(elements []int) {
	if len(elements) <= 10000 {
		sort.BubbleSort(elements)
		return
	}
	sort.Sort(elements)
}
