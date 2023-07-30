package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(e error)  {
	if e != nil {
		log.Fatalln(e)
	}
}

type MyStack []string

func (s *MyStack) Push(val string) {
	*s = append(*s, val)
}

func (s *MyStack) Pop() string {
	ret := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return ret
}

func (s *MyStack) Top() string {
	ret := (*s)[len(*s)-1]
	return ret
}

const DISK_SIZE = 70000000
const SPACE_REQUIRED = 30000000

func main() {
	file, err := os.Open("./day_07/input.txt")
	check(err)

	sc := bufio.NewScanner(file)
	lines := make([]string, 0)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	var result int
	result = part1(lines)
	fmt.Println(result)
	result = part2(lines)
	fmt.Println(result)
}

type MyString string

func (s MyString) rIndex(substring string) int {
	for i:=len(s)-1-len(substring);i>=0;i-- {
		if s[i:i+len(substring)] == MyString(substring) {
			return i
		}
	}
	return -1
}

func create_directory_map(lines []string) map[string]int {
	var directories = make(map[string]int)
	var cur MyString = "/home"
	for _, line := range lines {
		input := strings.Split(line, " ")
		if input[0]=="$" {
			// command
			if input[1]=="cd" {
				if input[2]=="/"{
					cur = "/home"
				}else if input[2]==".." {
					cur = cur[:cur.rIndex("/")]
				}else {
					cur = cur + "/" + MyString(input[2])
				}
			}else if input[1]=="ls" {
				continue
			}
		}else {
			// ls result
			if input[0]!="dir" {
				temp_path := cur
				for temp_path != "" {
					temp_size, err := strconv.Atoi(input[0])
					check(err)
					directories[string(temp_path)] += temp_size
					temp_path = temp_path[:temp_path.rIndex("/")]
				}
			}
		}
	}
	return directories
}

func part1(lines []string) int {
	directories := create_directory_map(lines)
	result := 0
	for _, size := range directories {
		if size < 100000 {
			result += size
		}
	}
	return result
}

func part2(lines []string) int {
	directories := create_directory_map(lines)
	var space_occupied = directories["/home"]
	var free_space = DISK_SIZE-space_occupied
	var extra_space_needed = SPACE_REQUIRED-free_space
	var directory_to_delete = space_occupied 
	for _, size := range directories {
		if size >= extra_space_needed && size < directory_to_delete {
			directory_to_delete = size
		}
	}
	return directory_to_delete
}
