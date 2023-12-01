package main

import (
	"math"
)

type Point struct {
	x int
	y int
}

func moveTail(tail, head Point) Point {
	dx, dy := head.x-tail.x, head.y-tail.y
	if math.Abs(float64(dx))<=1 && math.Abs(float64(dy))<=1 {
		return tail
	}
	if dx != 0 {
		dx = dx/int(math.Abs(float64(dx)))
	}
	if dy != 0 {
		dy = dy/int(math.Abs(float64(dy)))
	}
	return Point{tail.x+dx, tail.y+dy}
}
