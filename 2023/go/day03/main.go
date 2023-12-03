package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var test bool
	flag.BoolVar(&test, "test", false, "true if testing")
	flag.Parse()

	input_file, err := os.Open("./day03/input.txt")
	if test {
		input_file, err = os.Open("./day03/test.txt")
	}
	if err != nil {
		log.Fatalln(err)
	}
	defer input_file.Close()

	sc := bufio.NewScanner(input_file)
	lines := make([]string, 0)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	answer1 := part1(lines)
	fmt.Println(answer1)
	answer2 := part2(lines)
	fmt.Println(answer2)
}

var digits map[rune]bool = map[rune]bool{
	'0': true,
	'1': true,
	'2': true,
	'3': true,
	'4': true,
	'5': true,
	'6': true,
	'7': true,
	'8': true,
	'9': true,
}

func getNumber(row, col int, lines []string) (int, error) {
	r := col
	for r < len(lines[row]) && digits[rune(lines[row][r])] {
		r += 1
	}
	return strconv.Atoi(lines[row][col:r])
}

func part1(lines []string) int {
	res := 0
	numbers_start := make(map[[2]int]bool)
	for i, line := range lines {
		for j, c := range line {
			if digits[c] || c == '.' {
				continue
			}
			for _, rr := range []int{i - 1, i, i + 1} {
				for _, cc := range []int{j - 1, j, j + 1} {
					if rr == i && cc == j {
						continue
					}
					if digits[rune(lines[rr][cc])] {
						l := cc
						for l >= 0 && digits[rune(lines[rr][l])] {
							l -= 1
						}
						numbers_start[[2]int{rr, l + 1}] = true
					}
				}
			}
		}
	}
	for key := range numbers_start {
		num, _ := getNumber(key[0], key[1], lines)
		res += num
	}
	return res
}

func part2(lines []string) int {
	res := 0
	numbers_start := make(map[[2]int]map[[2]int]bool)
	for i, line := range lines {
		for j, c := range line {
			if c != '*' {
				continue
			}
			for _, rr := range []int{i - 1, i, i + 1} {
				for _, cc := range []int{j - 1, j, j + 1} {
					if rr == i && cc == j {
						continue
					}
					if digits[rune(lines[rr][cc])] {
						l := cc
						for l >= 0 && digits[rune(lines[rr][l])] {
							l -= 1
						}
						if numbers_start[[2]int{i, j}] == nil {
							numbers_start[[2]int{i, j}] = make(map[[2]int]bool)
						}
						numbers_start[[2]int{i, j}][[2]int{rr, l + 1}] = true
					}
				}
			}
		}
	}
	for _, value := range numbers_start {
		keys := make([][2]int, 0, len(value))
		for k := range value {
			keys = append(keys, k)
		}
		if len(keys) == 2 {
			product := 1
			for _, key := range keys {
				row, col := key[0], key[1]
				r := col
				for r < len(lines[row]) && digits[rune(lines[row][r])] {
					r += 1
				}
				num, _ := strconv.Atoi(lines[row][col:r])
				product *= num
			}
			res += product
		}
	}
	return res
}
