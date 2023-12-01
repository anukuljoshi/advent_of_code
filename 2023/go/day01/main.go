package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	input_file, err := os.Open("./day01/input.txt")
	if err!=nil {
		log.Fatalln(err)
	}
	defer input_file.Close()

	sc := bufio.NewScanner(input_file)
	lines := make([]string, 0)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	answer1 := Part1(lines)
	fmt.Println(answer1)
}

var digits map[string]int = map[string]int{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

func minIndex(nums []int) int {
	mn := nums[0]
	idx := 0
	for i, n := range nums {
		if n<mn {
			idx = i
			mn = n
		}
	}
	return idx
}

func maxIndex(nums []int) int {
	mx := nums[0]
	idx := 0
	for i, n := range nums {
		if n>mx {
			idx = i
			mx = n
		}
	}
	return idx
}

func Part1(lines []string) int {
	res := 0
	for _, line := range lines {
		first_indices := make([]int, 10)
		last_indices := make([]int, 10)
		for key, value := range digits {
			first_indices[value] = strings.Index(line, key)
			last_indices[value] = strings.LastIndex(line, key)
			if first_indices[value]==-1 {
				first_indices[value] = math.MaxInt32
			}
		}
		first, last := minIndex(first_indices), maxIndex(last_indices)
		number := first*10 + last
		res += number
	}
	return res
}
