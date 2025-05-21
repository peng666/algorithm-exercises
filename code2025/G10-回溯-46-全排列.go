package main

import "fmt"

func permute(nums []int) [][]int {
	var res [][]int
	var path []int
	visited := make(map[int]bool, len(nums))

	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			temp := make([]int, len(nums)) // 这里用make初始化长度不能为0，一定要够才能copy
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for _, num := range nums {
			if visited[num] {
				continue
			}
			visited[num] = true
			path = append(path, num)
			dfs()
			visited[num] = false
			path = path[:len(path)-1]
		}
	}
	dfs() // 记住要调用
	return res
}

func main() {
	res := permute([]int{1, 2, 3})
	fmt.Println(res)
}
