package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	queue := []*TreeNode{root}
	isNeedReverse := false
	for len(queue) > 0 {
		var temp []int
		var nextLevel []*TreeNode
		for _, node := range queue {
			temp = append(temp, node.Val)
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		if isNeedReverse {
			temp = reverseArray(temp)
		}
		res = append(res, temp)
		queue = nextLevel
		isNeedReverse = !isNeedReverse
	}
	return res
}

func reverseArray(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left += 1
		right -= 1
	}
	return nums
}

func main() {
	node15 := &TreeNode{Val: 15}
	node7 := &TreeNode{Val: 7}
	node20 := &TreeNode{Val: 20, Left: node15, Right: node7}
	node9 := &TreeNode{Val: 9}
	node3 := &TreeNode{Val: 3, Left: node9, Right: node20}
	res := zigzagLevelOrder(node3)
	fmt.Println(res)
}
