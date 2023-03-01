package main

import "sort"

func Ascending(n []int) []int {
	var result = n

	sort.Sort(sort.IntSlice(n))

	return result
}
