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
	flag.BoolVar(&test, "test", false, "use test file as input")
	flag.Parse()

	input_file, err := os.Open("./day04/input.txt")
	if test {
		input_file, err = os.Open("./day04/test.txt")
	}
	if err != nil {
		log.Fatalln(err)
	}
	lines := make([]string, 0)
	sc := bufio.NewScanner(input_file)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	answer1 := part1(lines)
	fmt.Println(answer1)
	answer2 := part2(lines)
	fmt.Println(answer2)
}

func part1(lines []string) int {
	res := 0
	for _, line := range lines {
		card := strings.Split(line, ":")
		numbers := strings.Split(card[1], "|")
		winning := strings.Split(numbers[0], " ")
		having := strings.Split(numbers[1], " ")
		winning_numbers := make(map[int]bool)
		for _, w := range winning {
			n, err := strconv.Atoi(strings.TrimSpace(w))
			if err != nil {
				continue
			}
			winning_numbers[n] = true
		}
		score := 0
		for _, h := range having {
			n, _ := strconv.Atoi(strings.TrimSpace(h))
			if winning_numbers[n] {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}
		res += score
	}
	return res
}

func part2(lines []string) int {
	res := 0
	copies := make(map[int]int)
	for i := range lines {
		copies[i] = 1
	}
	for i, line := range lines {
		for j := 0; j < copies[i]; j++ {
			card := strings.Split(line, ":")
			numbers := strings.Split(card[1], "|")
			winning := strings.Split(numbers[0], " ")
			having := strings.Split(numbers[1], " ")
			winning_numbers := make(map[int]bool)
			for _, w := range winning {
				n, err := strconv.Atoi(strings.TrimSpace(w))
				if err != nil {
					continue
				}
				winning_numbers[n] = true
			}
			match := 0
			for _, h := range having {
				n, _ := strconv.Atoi(strings.TrimSpace(h))
				if winning_numbers[n] {
					match += 1
				}
			}
			for c := i + 1; c < i+1+match; c++ {
				copies[c]++
			}

		}
	}
	for _, val := range copies {
		res += val
	}
	return res
}
