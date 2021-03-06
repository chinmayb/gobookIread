package main

import "sort"

// removeDuplicates return the number of unique values
func removeDuplicates(data []int) int {
	size, j := len(data), 0
	if size == 0 {
		return 0
	}

	sort.Ints(data)
	for i := 1; i < size; i++ {
		if data[i] != data[j] {
			j++
			data[j] = data[i]
		}
	}
	return j + 1
}
