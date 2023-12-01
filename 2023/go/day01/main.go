package main

import (
	"bufio"
	"fmt"
	"log"
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
	answer2 := Part2(lines)
	fmt.Println(answer1)
	fmt.Println(answer2)
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
			f := strings.Index(line, key)
			l := strings.LastIndex(line, key)
			if f==-1 {
				f = len(line)
			}
			first_indices[value] = f
			last_indices[value] = l
		}
		first, last := minIndex(first_indices), maxIndex(last_indices)
		number := first*10 + last
		res += number
	}
	return res
}

var spell map[string]int = map[string]int {
	"zero": 0,
	"one": 1,
	"two": 2,
	"three": 3,
	"four": 4,
	"five": 5,
	"six": 6,
	"seven": 7,
	"eight": 8,
	"nine": 9,
}

func Part2(lines []string) int {
	res := 0
	for _, line := range lines {
		first_indices := make([]int, 10)
		last_indices := make([]int, 10)
		for key, value := range digits {
			f := strings.Index(line, key)
			l := strings.LastIndex(line, key)
			if f==-1 {
				f = len(line)
			}
			first_indices[value] = f
			last_indices[value] = l
		}
		for key, value := range spell {
			f := strings.Index(line, key)
			l := strings.LastIndex(line, key)
			if f==-1 {
				f = len(line)
			}
			first_indices[value] = min(first_indices[value], f)
			last_indices[value] = max(last_indices[value], l)
		}
		first, last := minIndex(first_indices), maxIndex(last_indices)
		number := first*10 + last
		res += number
	}
	return res
}
