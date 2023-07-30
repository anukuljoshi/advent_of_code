package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func check(e error)  {
	if e != nil {
		log.Fatalln(e)
	}
}

func reverse(stack []byte) []byte {
	var rev = make([]byte, len(stack))
	for i, v := range stack {
		rev[len(rev)-1-i] = v
	}
	return rev
}

func parseStackInput(lines []string) [][]byte {
	var stacks [][]byte
	for _, line := range lines {
		if line == "" {
			break
		}
		stack := 0
		for {
			if stack > len(stacks)-1 {
				stacks = append(stacks, []byte(nil))
			}
			if line[0] == '[' {
				stacks[stack] = append(stacks[stack], line[1])
			}
			if line = line[3:]; len(line) == 0 {
				break
			}
			line = line[1:] // Consume space
			stack++
		}
	}
	return stacks
}

func part1(lines []string) string {
	var res string
	stacks := parseStackInput(lines)
	for i:=0;i<len(stacks);i++ {
		stacks[i] = reverse(stacks[i])
	}
	for _, line := range lines {
		var quantity, from, to int
		if strings.HasPrefix(line, "move") {
			fmt.Sscanf(line, "move %d from %d to %d", &quantity,  &from, &to)
			from -= 1
			to -= 1
			for i:=0; i<quantity;i++ {
				item  := stacks[from][len(stacks[from])-1]
				stacks[from] = stacks[from][:len(stacks[from])-1]
				stacks[to] = append(stacks[to], item)
			}
		}
	}
	for _, stack := range stacks {
		res += string(stack[len(stack)-1])
	}
	return res
}

func part2(lines []string) string {
	var res string
	stacks := parseStackInput(lines)
	for i:=0;i<len(stacks);i++ {
		stacks[i] = reverse(stacks[i])
	}
	for _, line := range lines {
		var quantity, from, to int
		if strings.HasPrefix(line, "move") {
			fmt.Sscanf(line, "move %d from %d to %d", &quantity,  &from, &to)
			from -= 1
			to -= 1
			items := stacks[from][len(stacks[from])-quantity:]
			stacks[from] = stacks[from][:len(stacks[from])-quantity]
			stacks[to] = append(stacks[to], items...)
		}
	}
	for _, stack := range stacks {
		res += string(stack[len(stack)-1])
	}
	return res
}

func main() {
	file, err := os.Open("./day_05/input.txt")
	check(err)
	sc := bufio.NewScanner(file)
	lines := make([]string, 0)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	var res string
	res = part1(lines)
	fmt.Println(res)
	res = part2(lines)
	fmt.Println(res)
}