package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var test bool
	flag.BoolVar(&test, "test", false, "use test file")
	flag.Parse()

	file_path := "./day07/input.txt"
	if test {
		file_path = "./day07/test.txt"
	}
	input_file, err := os.Open(file_path)
	if err != nil {
		log.Fatalln(err)
	}
	defer input_file.Close()
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

const (
	FIVE     = 6
	FOUR     = 5
	FULL     = 4
	THREE    = 3
	TWO_PAIR = 2
	ONE_PAIR = 1
	HIGH     = 0
)

func cmpFunc(cardOrder map[rune]int, a, b string) int {
	n := len(a)
	for i := 0; i < n; i++ {
		ca, cb := rune(a[i]), rune(b[i])
		o1, _ := cardOrder[ca]
		o2, _ := cardOrder[cb]
		if o1 != o2 {
			return o1 - o2
		}
	}
	return 0
}

func part1(lines []string) int {
	var cardOrder map[rune]int = map[rune]int{
		'2': 0,
		'3': 1,
		'4': 2,
		'5': 3,
		'6': 4,
		'7': 5,
		'8': 6,
		'9': 7,
		'T': 8,
		'J': 9,
		'Q': 10,
		'K': 11,
		'A': 12,
	}

	res := 0
	ranks := make(map[int][]string, 7)
	bids := make(map[string]int)
	for _, line := range lines {
		hand := strings.Split(line, " ")
		n, _ := strconv.Atoi(hand[1])
		bids[hand[0]] = n
		counts := make([]int, len(cardOrder))
		for _, c := range hand[0] {
			counts[cardOrder[c]] += 1
		}
		slices.SortFunc[[]int](counts, func(a, b int) int {
			return b - a
		})
		if counts[0] == 5 {
			// five of a kind
			ranks[FIVE] = append(ranks[FIVE], hand[0])
			slices.SortFunc[[]string](ranks[FIVE], func(a, b string) int {
				return cmpFunc(cardOrder, a, b)
			})
		} else if counts[0] == 4 {
			// four of a kind
			ranks[FOUR] = append(ranks[FOUR], hand[0])
			slices.SortFunc[[]string](ranks[FOUR], func(a, b string) int {
				return cmpFunc(cardOrder, a, b)
			})
		} else if counts[0] == 3 && counts[1] == 2 {
			// full house
			ranks[FULL] = append(ranks[FULL], hand[0])
			slices.SortFunc[[]string](ranks[FULL], func(a, b string) int {
				return cmpFunc(cardOrder, a, b)
			})
		} else if counts[0] == 3 {
			// three of a kind
			ranks[THREE] = append(ranks[THREE], hand[0])
			slices.SortFunc[[]string](ranks[THREE], func(a, b string) int {
				return cmpFunc(cardOrder, a, b)
			})
		} else if counts[0] == 2 && counts[1] == 2 {
			// two pairs
			ranks[TWO_PAIR] = append(ranks[TWO_PAIR], hand[0])
			slices.SortFunc[[]string](ranks[TWO_PAIR], func(a, b string) int {
				return cmpFunc(cardOrder, a, b)
			})
		} else if counts[0] == 2 {
			// pair
			ranks[ONE_PAIR] = append(ranks[ONE_PAIR], hand[0])
			slices.SortFunc[[]string](ranks[ONE_PAIR], func(a, b string) int {
				return cmpFunc(cardOrder, a, b)
			})
		} else if counts[0] == 1 {
			// high
			ranks[HIGH] = append(ranks[HIGH], hand[0])
			slices.SortFunc[[]string](ranks[HIGH], func(a, b string) int {
				return cmpFunc(cardOrder, a, b)
			})
		}
	}
	r := 1
	for i := 0; i < 7; i++ {
		for _, hand := range ranks[i] {
			res += (r * bids[hand])
			r += 1
		}
	}
	return res
}

func part2(lines []string) int {
	var cardOrder map[rune]int = map[rune]int{
		'J': 0,
		'2': 1,
		'3': 2,
		'4': 3,
		'5': 4,
		'6': 5,
		'7': 6,
		'8': 7,
		'9': 8,
		'T': 9,
		'Q': 10,
		'K': 11,
		'A': 12,
	}

	res := 0
	ranks := make(map[int][]string, 7)
	bids := make(map[string]int)
	for _, line := range lines {
		hand := strings.Split(line, " ")
		n, _ := strconv.Atoi(hand[1])
		bids[hand[0]] = n
		counts := make([]int, len(cardOrder))
		for _, c := range hand[0] {
			counts[cardOrder[c]] += 1
		}
		slices.SortFunc[[]int](counts[1:], func(a, b int) int {
			return a - b
		})
		counts[len(counts)-1] += counts[0]
		counts[0] = 0
		slices.SortFunc[[]int](counts, func(a, b int) int {
			return b - a
		})
		if counts[0] == 5 {
			// five of a kind
			ranks[FIVE] = append(ranks[FIVE], hand[0])
			slices.SortFunc[[]string](ranks[FIVE], func(a, b string) int {
				return cmpFunc(cardOrder, a, b)
			})
		} else if counts[0] == 4 {
			// four of a kind
			ranks[FOUR] = append(ranks[FOUR], hand[0])
			slices.SortFunc[[]string](ranks[FOUR], func(a, b string) int {
				return cmpFunc(cardOrder, a, b)
			})
		} else if counts[0] == 3 && counts[1] == 2 {
			// full house
			ranks[FULL] = append(ranks[FULL], hand[0])
			slices.SortFunc[[]string](ranks[FULL], func(a, b string) int {
				return cmpFunc(cardOrder, a, b)
			})
		} else if counts[0] == 3 {
			// three of a kind
			ranks[THREE] = append(ranks[THREE], hand[0])
			slices.SortFunc[[]string](ranks[THREE], func(a, b string) int {
				return cmpFunc(cardOrder, a, b)
			})
		} else if counts[0] == 2 && counts[1] == 2 {
			// two pairs
			ranks[TWO_PAIR] = append(ranks[TWO_PAIR], hand[0])
			slices.SortFunc[[]string](ranks[TWO_PAIR], func(a, b string) int {
				return cmpFunc(cardOrder, a, b)
			})
		} else if counts[0] == 2 {
			// pair
			ranks[ONE_PAIR] = append(ranks[ONE_PAIR], hand[0])
			slices.SortFunc[[]string](ranks[ONE_PAIR], func(a, b string) int {
				return cmpFunc(cardOrder, a, b)
			})
		} else if counts[0] == 1 {
			// high
			ranks[HIGH] = append(ranks[HIGH], hand[0])
			slices.SortFunc[[]string](ranks[HIGH], func(a, b string) int {
				return cmpFunc(cardOrder, a, b)
			})
		}
	}
	r := 1
	for i := 0; i < 7; i++ {
		for _, hand := range ranks[i] {
			res += (r * bids[hand])
			r += 1
		}
	}
	return res
}
