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

func part1(lines []string) int  {
	
	res := 0
	for _, sack := range lines {
		first := make([]int, 53)
		second := make([]int, 53)
		for i, j := 0, len(sack)/2; j<len(sack); i, j = i+1, j+1 {
			// lower half
			var f, s int
			if int(sack[i]) <= int('Z'){
				// uppercase
				f = int(sack[i])-int('A')+1+26
				first[f] += 1
			}else if int(sack[i]) <= int('z') {
				// lowercase
				f = int(sack[i])-int('a')+1
				first[f] += 1
			}
			// upper half
			if int(sack[j]) <= int('Z'){
				// uppercase
				s = int(sack[j])-int('A')+1+26
				second[s] += 1
			}else if int(sack[j]) <= int('z') {
				// lowercase
				s = int(sack[j])-int('a')+1
				second[s] += 1
			}
		}
		for i := range(first) {
			if first[i] > 0 && second[i] > 0 {
				res += i
			}
		}
	}
		return res
}

func part2(lines []string) int {
	i := 0
	res := 0
	for i < len(lines) {
		in1 := lines[i]
		in2 := lines[i+1]
		in3 := lines[i+2]

		var map1 = make([]int, 256)
		var map2 = make([]int, 256)
		var map3 = make([]int, 256)
		for _, v := range(in1) {
			map1[v] += 1
		}
		for _, v := range(in2) {
			map2[v] += 1
		}
		for _, v := range(in3) {
			map3[v] += 1
		}
		for i:=0; i<256; i++ {
			if map1[i]>0 && map2[i]>0 && map3[i]>0 {
				if i < 95 {
					res += i-64+26
				}else {
					res += i-96
				}
			}
		}
		i += 3
	}
	return res
}

func main() {
	file, err := os.Open("./day_03/input.txt")
	check(err)
	defer file.Close()

	sc := bufio.NewScanner(file)
	lines := make([]string, 0)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	res := part1(lines)
	fmt.Println(res)
	res = part2(lines)
	fmt.Println(res)
}
