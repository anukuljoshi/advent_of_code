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
					if rr < 0 || rr >= len(lines) || cc < 0 || cc >= len(line) || !digits[rune(lines[rr][cc])] {
						continue
					}
					for cc > 0 && digits[rune(lines[rr][cc-1])] {
						cc -= 1
					}
					numbers_start[[2]int{rr, cc}] = true
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
	for i, line := range lines {
		for j, c := range line {
			if c != '*' {
				continue
			}
			numbers_start := make(map[[2]int]bool)
			for _, rr := range []int{i - 1, i, i + 1} {
				for _, cc := range []int{j - 1, j, j + 1} {
					if rr < 0 || rr >= len(lines) || cc < 0 || cc >= len(line) || !digits[rune(lines[rr][cc])] {
						continue
					}
					for cc > 0 && digits[rune(lines[rr][cc-1])] {
						cc -= 1
					}
					numbers_start[[2]int{rr, cc}] = true
				}
			}
			keys := make([][2]int, 0, len(numbers_start))
			for k := range numbers_start {
				keys = append(keys, k)
			}
			if len(keys) != 2 {
				continue
			}
			product := 1
			for key := range numbers_start {
				num, _ := getNumber(key[0], key[1], lines)
				product *= num
			}
			res += product
		}
	}
	return res
}
