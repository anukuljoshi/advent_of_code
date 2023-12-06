package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var test bool
	flag.BoolVar(&test, "test", false, "use test file input")
	flag.Parse()

	filename := "./day06/input.txt"
	if test {
		filename = "./day06/test.txt"
	}
	input_file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	lines := make([]string, 0)
	sc := bufio.NewScanner(input_file)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	answer1 := part1(lines)
	fmt.Println("Part 1:", answer1)
	answer2 := part2(lines)
	fmt.Println("Part 2", answer2)
}

func convertToIntArray(s []string) []int {
	res := make([]int, 0)
	for _, c := range s {
		n, err := strconv.Atoi(c)
		if err != nil {
			continue
		}
		res = append(res, n)
	}
	return res
}

func binarySearch(num, target int) int {
	l, r := 0, int(num/2)
	for l < r {
		m := l + int((r-l)/2)
		if m*(num-m) > target {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func part1(lines []string) int {
	res := 1
	time := convertToIntArray(strings.Split(strings.Split(lines[0], ":")[1], " "))
	distance := convertToIntArray(strings.Split(strings.Split(lines[1], ":")[1], " "))
	for i := 0; i < len(time); i++ {
		t := time[i]
		d := distance[i]
		p := binarySearch(t, d)
		ways := (t + 1) - p*2
		res *= ways
	}
	return res
}

func convertToString(s string) string {
	res := ""
	for _, c := range s {
		_, err := strconv.Atoi(string(c))
		if err != nil {
			continue
		}
		res += string(c)
	}
	return res
}

func part2(lines []string) int {
	t, _ := strconv.Atoi(convertToString(strings.Split(lines[0], ":")[1]))
	d, _ := strconv.Atoi(convertToString(strings.Split(lines[1], ":")[1]))
	p := binarySearch(t, d)
	ways := (t + 1) - p*2
	return ways
}
