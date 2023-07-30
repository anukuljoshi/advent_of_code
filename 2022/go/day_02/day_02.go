package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type ListNode struct {
	next *ListNode
	prev *ListNode
	val int
}

var rock = &ListNode{
	next: nil,
	prev: nil,
	val: 1,
}
var paper = &ListNode{
	next: nil,
	prev: nil,
	val: 2,
}
var scissors = &ListNode{
	next: nil,
	prev: nil,
	val: 3,
}

var codeABC = map[string]*ListNode {
	"A": rock, 
	"B": paper,
	"C": scissors,
}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

var codeXYZ = map[string]*ListNode {
	"X": rock, 
	"Y": paper,
	"Z": scissors,
}

func Winner(op, me string) string {
	rock.next = paper
	paper.next = scissors
	scissors.next = rock

	if codeABC[op] == codeXYZ[me] {
		// draw
		return "D"
	}
	if codeABC[op].next == codeXYZ[me] {
		// win
		return "W"
	}
	return "L"
}

func part1(lines []string) int {
	var res int
	var resultScore =  make(map[string]int)
	resultScore["L"] = 0
	resultScore["D"] = 3
	resultScore["W"] = 6

	for _, line := range lines {
		chances := strings.Split(line, " ")
		abc, xyz := string(chances[0]), string(chances[1])
		w := Winner(abc, xyz)
		res += resultScore[w]
		res += codeXYZ[xyz].val
	}
	return res
}

func ShapeScore(op, result string) int {
	rock.next = paper
	paper.prev = rock	
	paper.next = scissors
	scissors.prev = paper
	scissors.next = rock
	rock.prev = scissors

	if result == "X" {
		// lose
		return codeABC[op].prev.val
	}
	if result == "Y" {
		// draw
		return codeABC[op].val
	}
	// win
	return codeABC[op].next.val
}

func part2(lines []string) int {
	var res int
	var resultScore =  make(map[string]int)
	resultScore["X"] = 0
	resultScore["Y"] = 3
	resultScore["Z"] = 6
	for _, line := range lines {
		chances := strings.Split(line, " ")
		abc, xyz := string(chances[0]), string(chances[1])
		w := ShapeScore(abc, xyz)
		res += resultScore[xyz]
		res += w
	}
	return res
}

func main() {
	file, err := os.Open("./day_02/input.txt")
	check(err)
	defer file.Close()

	sc := bufio.NewScanner(file)
	lines := make([]string, 0)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	var res int
	res = part1(lines)
	fmt.Println("Part 1: ", res)
	res = part2(lines)
	fmt.Println("Part 2: ", res)
}