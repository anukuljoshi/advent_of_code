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
    file, err := os.Open("./day_08/input.txt")
    check(err)
        
    sc := bufio.NewScanner(file)
    lines := make([]string, 0)
    for sc.Scan() {
        lines = append(lines, sc.Text())
    }
    grid := parse_input(lines)
    var result int
    result = part1(grid)
    fmt.Println(result)
	result = part2(grid)
    fmt.Println(result)
}
