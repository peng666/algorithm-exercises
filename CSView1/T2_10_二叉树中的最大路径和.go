package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ans := math.MinInt32        // 弄成全局变量
	var dfs func(*TreeNode) int // 定义一个函数
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0 // 递归终止条件：就是遇到叶子节点的空指针，为0
		}

		leftSum := dfs(root.Left)
		rightSum := dfs(root.Right)
		rootSum := root.Val + leftSum + rightSum // 这就是当前最优路径
		ans = max(ans, rootSum)
		return max(max(leftSum, rightSum)+root.Val, 0) // 如果是负数就舍弃该路径
	}

	dfs(root)
	return ans
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func main() {
	node1 := &TreeNode{}
	res := maxPathSum(node1)
	fmt.Println(res)
}
