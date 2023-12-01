package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func check(err error) {
    if err != nil {
        log.Fatalln(err)
    }
}
func main() {
    file, err := os.Open("./day_09/input.txt")
    check(err)
    defer file.Close()

    sc := bufio.NewScanner(file)
    lines := make([]string, 0)
    for sc.Scan() {
        lines = append(lines, sc.Text())
    }
    moves := parseInput(lines)
    var res int
    res = part1(moves)
    fmt.Println(res)
    res = part2(moves)
    fmt.Println(res)
}
