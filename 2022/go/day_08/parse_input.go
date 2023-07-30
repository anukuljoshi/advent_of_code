package main

import (
	"strconv"
)

func stringToIntArray(s string) []int {
	var intArray = make([]int, 0)
	for _, c := range s {
		i, _ := strconv.Atoi(string(c))
		intArray = append(intArray, i)
	}
	return intArray
}

func parse_input(lines []string) [][]int {
	var heights = make([][]int, 0)
	for _, line := range lines {
		line_array := stringToIntArray(line)
		heights = append(heights, line_array)
	}
	return heights
}
