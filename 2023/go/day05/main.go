package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var test bool
	flag.BoolVar(&test, "test", false, "use test input")
	flag.Parse()

	input_file, err := os.Open("./day05/input.txt")
	if test {
		input_file, err = os.Open("./day05/test.txt")
	}
	if err != nil {
		log.Fatalln(err)
	}
	lines := make([]string, 0)
	sc := bufio.NewScanner(input_file)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	lines = append(lines, "")
	answer1 := part1(lines)
	fmt.Println(answer1)
	answer2 := part2(lines)
	fmt.Println(answer2)
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

// func part1(lines []string) int {
// 	res := math.MaxInt32
// 	conversion := make([]map[string]int, 0)
// 	seeds := convertToIntArray(strings.Split(strings.Split(lines[0], ":")[1], " "))
// 	for _, seed := range seeds {
// 		conversion = append(conversion, map[string]int{})
// 		key := "seed"
// 		value := seed
// 		conversion[len(conversion)-1][key] = value
// 		for i := 1; i < len(lines); i++ {
// 			line := lines[i]
// 			if len(line) == 0 {
// 				continue
// 			}
// 			if strings.Index(line, "-") >= 0 {
// 				key = strings.Split(strings.Split(line, " ")[0], "-")[2]
// 				conversion[len(conversion)-1][key] = value
// 				continue
// 			}
// 			m := convertToIntArray(strings.Split(line, " "))
// 			if value >= m[1] && value < m[1]+m[2] {
// 				value = m[0] + (value - m[1])
// 				conversion[len(conversion)-1][key] = value
// 				for i < len(lines) && len(lines[i]) != 0 {
// 					i++
// 				}
// 			}
// 		}
// 	}
// 	for _, convertor := range conversion {
// 		res = min(res, convertor["location"])
// 	}
// 	return res
// }

func part1(lines []string) int {
	res := math.MaxInt32
	convertors := make([]*Convertor, 0)
	var con *Convertor
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		if len(line) == 0 {
			if con != nil {
				convertors = append(convertors, con)
			}
			con = &Convertor{}
			continue
		}
		if strings.Index(line, "-") >= 0 {
			keys := strings.Split(strings.Split(line, " ")[0], "-")
			con.from = keys[0]
			con.to = keys[2]
			continue
		}
		m := convertToIntArray(strings.Split(line, " "))
		con.maps = append(con.maps, m)
	}

	// for _, c := range convertors {
	// 	fmt.Print(c, ", ")
	// }
	// fmt.Println()

	seeds := convertToIntArray(strings.Split(strings.Split(lines[0], ":")[1], " "))
	for _, seed := range seeds {
		value := seed
		for _, convertor := range convertors {
			for _, m := range convertor.maps {
				if m.Contains(value) {
					value = m.Get(value)
					break
				} else {
					value = m.Get(value)
				}
			}
		}
		res = min(res, value)
	}
	return res
}

type Map []int

func (m Map) Contains(value int) bool {
	if value >= m[1] && value < m[1]+m[2] {
		return true
	}
	return false
}

func (m Map) Get(key int) int {
	value := key
	if m.Contains(value) {
		value = m[0] + (value - m[1])
	}
	return value
}

type Convertor struct {
	from string
	to   string
	maps []Map
}

func part2(lines []string) int {
	res := math.MaxInt64
	convertors := make([]*Convertor, 0)
	var con *Convertor
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		if len(line) == 0 {
			if con != nil {
				convertors = append(convertors, con)
			}
			con = &Convertor{}
			continue
		}
		if strings.Index(line, "-") >= 0 {
			keys := strings.Split(strings.Split(line, " ")[0], "-")
			con.from = keys[0]
			con.to = keys[2]
			continue
		}
		m := convertToIntArray(strings.Split(line, " "))
		con.maps = append(con.maps, m)
	}
	seeds := convertToIntArray(strings.Split(strings.Split(lines[0], ":")[1], " "))
	for i := 0; i < len(seeds); i += 2 {
		for seed := seeds[i]; seed < seeds[i]+seeds[i+1]; seed++ {
			value := seed
			for _, convertor := range convertors {
				for _, m := range convertor.maps {
					if m.Contains(value) {
						value = m.Get(value)
						break
					} else {
						value = m.Get(value)
					}
				}
			}
			res = min(res, value)
		}
	}
	return res
}
