package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func min(a, b int) int {
	if a<b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func part1(lines []string) int {
	var count int
	var first [2]int
	var second [2]int
	for _, line := range lines {
		fmt.Sscanf(line, "%d-%d,%d-%d", &first[0], &first[1], &second[0], &second[1])
		result_range := [2]int{min(first[0], second[0]), max(first[1], second[1])}
		if result_range==first || result_range==second {
			count++
		}
	}
	return count
}

func part2(lines []string) int {
	var count int
	var first [2]int
	var second [2]int
	for _, line := range lines {
		fmt.Sscanf(line, "%d-%d,%d-%d", &first[0], &first[1], &second[0], &second[1])
		if second[0]<=first[1] && second[1]>=first[0] {
			count++
		}
	}
	return count
}

func main()  {
	file, err := os.Open("./day_04/input.txt")
	check(err)
	defer file.Close()
	
	sc := bufio.NewScanner(file)
	var lines []string
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	// answer to  part 1
	var res int
	res = part1(lines)
	fmt.Println(res)
	res = part2(lines)
	fmt.Println(res)
}
