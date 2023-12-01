package main

func part2(moves []Move) int {
	var res int
	rope := make([]Point, 10)
	var visited = make(map[Point]bool)
	visited[rope[9]] = true
	for _, move := range moves {
		for s:=0;s<move.step;s++ {
			switch move.direction {
				case "L":
					rope[0].x -= 1
				case "R":
					rope[0].x += 1
				case "U":
					rope[0].y += 1
				case "D":
					rope[0].y -= 1
			}
			for t:=0;t<len(rope)-1;t++ {
				rope[t+1] = moveTail(rope[t+1], rope[t])
			}
			visited[rope[9]] = true
		}
	}
	res = len(visited)
	return res
}
