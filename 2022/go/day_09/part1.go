package main

func part1(moves []Move) int {
	head := Point{0, 0}
	tail := Point{0, 0}
	var visited = make(map[Point]bool)

	for _, move := range moves {
		for s:=0;s<move.step;s++ {
			switch move.direction {
			case "L":
				head.x -= 1
			case "R":
				head.x += 1
			case "U":
				head.y += 1
			case "D":
				head.y -= 1
			}
			tail = moveTail(tail, head)
			visited[tail] = true
		}
	}
	return len(visited)
}
