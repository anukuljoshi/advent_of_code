package main

func traverseRow(grid [][]int, visible *[][]int, row int) {
    var maxHeight = grid[row][0]
    for j:=1;j<len(grid[0])-1;j++ {
        if grid[row][j] > maxHeight {
            maxHeight = grid[row][j]
            (*visible)[row][j] = 1
        }
    }
}

func traverseRowReverse(grid [][]int, visible *[][]int, row int) {
    var maxHeight = grid[row][len(grid[0])-1]
    for j:=len(grid[0])-2;j>0;j-- {
        if grid[row][j] > maxHeight {
            maxHeight = grid[row][j]
            (*visible)[row][j] = 1
        }
    }
}

func traverseCol(grid [][]int, visible *[][]int, col int) {
    var maxHeight = grid[0][col]
    for i:=1;i<len(grid)-1;i++ {
        if grid[i][col] > maxHeight {
            maxHeight = grid[i][col]
            (*visible)[i][col] = 1
        }
    }
}

func traverseColReverse(grid [][]int, visible *[][]int, col int) {
    var maxHeight = grid[len(grid)-1][col]
    for i:=len(grid)-2;i>0;i-- {
        if grid[i][col] > maxHeight {
            maxHeight = grid[i][col]
            (*visible)[i][col] = 1
        }
    }
}

func countVisibleTrees(visible [][]int) int {
    var res = 0
    for i:=0;i<len(visible);i++ {
        for j:=0;j<len(visible[0]);j++ {
            res += visible[i][j]
        }
    }
    return res
}

func part1(grid [][]int) int {
    var res int
    var m, n = len(grid), len(grid[0])
    var visible = make([][]int, m)
    for i := range visible {
        visible[i] = make([]int, n)
    }

    // first row
    for j:=1;j<len(grid[0])-1;j++ {
        traverseCol(grid, &visible, j)
    }

    // first col
    for i:=1;i<len(grid)-1;i++ {
        traverseRow(grid, &visible, i)
    }

    // last row
    for j:=1;j<len(grid[0])-1;j++ {
        traverseColReverse(grid, &visible, j)
    }

    // last col
    for i:=1;i<len(grid)-1;i++ {
        traverseRowReverse(grid, &visible, i)
    }

    res = countVisibleTrees(visible)
    // add perimeter trees to result
    res += (2*(m+n))-4
    return res
}
