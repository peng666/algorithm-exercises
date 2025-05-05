package CSView1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 也就是Z字形
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	ans := [][]int{}
	flag := false
	for len(queue) > 0 {
		tmp := []int{}
		q := []*TreeNode{}

		for _, node := range queue { // 遍历这个切片就行啦。不用出队啥的
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		if flag {
			tmp = reverse(tmp)
		}

		ans = append(ans, tmp)
		queue = q
		flag = !flag
	}
	return ans
}

func reverse(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++ // 记得移动指针呀~~老是忘记
		right--
	}
	return nums
}
