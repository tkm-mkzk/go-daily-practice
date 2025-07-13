package main

import "sort"

func MergeSortedSlices(slice1, slice2 []int) []int {
	if len(slice1) == 0 && len(slice2) == 0 {
		return []int{}
	} else if len(slice1) == 0 {
		return slice2
	} else if len(slice2) == 0 {
		return slice1
	}

	newSlice := make([]int, 0, len(slice1)+len(slice2))

	newSlice = append(newSlice, slice1...)
	newSlice = append(newSlice, slice2...)
	sort.Ints(newSlice)

	return newSlice
}
