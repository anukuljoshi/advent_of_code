package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func part1(lines []string) int  {
	var res int = 0
	var cur int
	for _, line := range(lines) {
		if line != "" {
			temp, err := strconv.Atoi(line) 
			check(err)
			cur += temp
		} else {
			res = max(res, cur)
			cur = 0
		}
	}
	res = max(res, cur)
	return res
}

func sum(s []int) int  {
	ret := 0
	for _, v := range(s) {
		ret += v
	}
	return ret
}

func part2(lines []string) int  {
	res := []int{}
	var cur int
	for _, line := range(lines) {
		if line != "" {
			temp, err := strconv.Atoi(line) 
			check(err)
			cur += temp
		} else {
			res = append(res, cur)
			cur = 0
		}
	}
	res = append(res, cur)
	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	return sum(res[len(res)-3:])
}

func main() {
	file, err := os.Open("./day_01/input.txt")
	check(err)
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		lines = append(lines, scanner.Text())
	}
	var res int
	res = part1(lines)
	fmt.Println(res)
	res = part2(lines)
	fmt.Println(res)
}
