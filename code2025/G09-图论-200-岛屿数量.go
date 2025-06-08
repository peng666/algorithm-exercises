package main

import "fmt"

func numIsIands(grid [][]byte) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				res += 1
				dfs(grid, i, j)
			}
		}
	}
	return res
}

func dfs(grid [][]byte, r int, c int) {
	//边界条件，退出
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
		return
	}

	// 判断有没有遍历过(1表示需要遍历且还没遍历过。2表示遍历过，0表示不需要遍历)
	if grid[r][c] != '1' {
		return
	}

	// 第1次遍历到这里，标记为2
	grid[r][c] = '2'

	// 4个方向继续dfs递归标记
	dfs(grid, r-1, c)
	dfs(grid, r+1, c)
	dfs(grid, r, c-1)
	dfs(grid, r, c+1)
}

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	res := numIsIands(grid)
	fmt.Println(res)
}
