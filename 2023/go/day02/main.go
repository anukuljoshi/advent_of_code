package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input_file, err := os.Open("./day02/input.txt")
	if err!=nil {
		log.Fatalln(err)
	}
	sc := bufio.NewScanner(input_file)
	games := make([][]string, 0)
	for sc.Scan() {
		line := sc.Text()
		a := strings.Split(line, ":")
		b := strings.Split(a[1], ";")
		games = append(games, b)
	}
	answer1 :=  part1(games)
	fmt.Println(answer1)
}

var available map[string]int = map[string]int {
	"red": 12,
	"green": 13,
	"blue": 14,
}

// part1
func part1(games [][]string) int {
	res := 0
	for i, game := range games {
		count := make([]map[string]int, 0)
		for _, g := range game {
			cur := make(map[string]int)
			cubes := strings.Split(g, ",")
			for _, cube := range cubes {
				color := strings.Split(strings.Trim(cube, " "), " ")
				cur[color[1]], _ = strconv.Atoi(color[0])
			}
			count = append(count, cur)
		}
		possible := true
		for _, c := range count {
			for key := range c {
				if c[key]>available[key] {
					possible = false
				}
			}
		}
		if possible {
			res += (i+1)
		}
	}
	return res
}
