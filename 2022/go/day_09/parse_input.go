package main

import (
	"strconv"
	"strings"
)

type Move struct {
    direction string
    step int
}


func parseInput(lines []string) []Move {
    var moves = make([]Move, 0)
    for _, line := range lines {
        input := strings.Split(line, " ")
        direction := input[0]
        step, err := strconv.Atoi(input[1])
        check(err)
        moves = append(moves, Move{direction: direction, step: step})
    }
    return moves
}
