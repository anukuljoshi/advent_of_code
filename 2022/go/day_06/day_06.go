package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func check(e error)  {
	if e != nil {
		log.Fatalln(e)
	}
}

func part1(lines []string) int {
	var freq = make([]int, 256)
	for _, line := range lines {
		freq[line[0]] += 1
		freq[line[1]] += 1
		freq[line[2]] += 1
		freq[line[3]] += 1
		var i, j = 0, 3
		for j < len(line){
			var dup = 0
			for _, f := range freq {
				if f > 1 {
					dup = 1
				}
			}
			if dup == 0 {
				return j+1
			}
			freq[line[i]] -= 1
			i += 1
			j += 1
			if j < len(line) {
				freq[line[j]] += 1
			}
		}
	}
	return -1
}

func part2(lines []string) int {
	var freq = make([]int, 256)
	for _, line := range lines {
		for i:=0;i<14;i++ {
			freq[line[i]] += 1
		}
		var i, j = 0, 13
		for j < len(line){
			var dup = 0
			for _, f := range freq {
				if f > 1 {
					dup = 1
				}
			}
			if dup == 0 {
				return j+1
			}
			freq[line[i]] -= 1
			i += 1
			j += 1
			if j < len(line) {
				freq[line[j]] += 1
			}
		}
	}
	return -1
}

func main() {
	file, err := os.Open("./day_06/input.txt")
	check(err)

	sc := bufio.NewScanner(file)
	var lines []string
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	var res int
	res = part1(lines)
	fmt.Println(res)
	res = part2(lines)
	fmt.Println(res)
}
