package main

import "fmt"

func generateParenthesis(n int) []string {
	var res []string
	var dfs func(string, int, int)
	dfs = func(paths string, left int, right int) {
		if left > n || left < right {
			return
		}

		if len(paths) == n*2 {
			res = append(res, paths)
			return
		}

		dfs(paths+"(", left+1, right)
		dfs(paths+")", left, right+1)
	}

	dfs("", 0, 0)
	return res
}

func main() {
	res := generateParenthesis(3)
	fmt.Println(res)
}
