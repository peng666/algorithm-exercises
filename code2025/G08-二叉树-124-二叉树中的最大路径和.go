package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	res := math.MinInt32
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftSum := dfs(root.Left)
		rightSum := dfs(root.Right)
		curSum := root.Val + leftSum + rightSum
		res = max(res, curSum)
		return max(0, root.Val+max(leftSum, rightSum))
	}
	dfs(root)
	return res
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func main() {
	node15 := &TreeNode{Val: 15}
	node7 := &TreeNode{Val: 7}
	node20 := &TreeNode{Val: 20, Left: node15, Right: node7}
	node9 := &TreeNode{Val: 9}
	node010 := &TreeNode{Val: -10, Left: node9, Right: node20}

	res := maxPathSum(node010)
	fmt.Println(res)

}
