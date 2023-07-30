package main

func traverseColDown(grid [][]int, row, col int) int {
	var visible = 0
	for i:=row+1;i<len(grid);i++ {
		visible++
		if grid[i][col] >= grid[row][col] {
			break
		}
	}
	return visible
}

func traverseColUp(grid [][]int, row, col int) int {
	var visible = 0
	for i:=row-1;i>=0;i-- {
		visible++
		if grid[i][col] >= grid[row][col] {
			break
		}
	}
	return visible
}

func traverseRowRight(grid [][]int, row, col int) int {
	var visible = 0
	for j:=col+1;j<len(grid[row]);j++ {
		visible++
		if grid[row][j] >= grid[row][col] {
			break
		}
	}
	return visible
}

func traverseRowLeft(grid [][]int, row, col int) int {
	var visible = 0
	for j:=col-1;j>=0;j-- {
		visible++
		if grid[row][j] >= grid[row][col] {
			break
		}
	}
	return visible
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func part2(grid [][]int) int {
    var res int
    var m, n = len(grid), len(grid[0])
    var visible = make([][]int, m)
    for i := range visible {
        visible[i] = make([]int, n)
    }

	for i:=0;i<m;i++ {
		for j:=0;j<n;j++ {
			res = max(res, traverseRowRight(grid, i, j) * traverseRowLeft(grid, i, j) * traverseColDown(grid, i, j) * traverseColUp(grid, i, j))
		}
	}
    return res
}
