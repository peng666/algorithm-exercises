package CSView1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	ans := []int{}
	for len(queue) > 0 {
		q := []*TreeNode{}
		for idx, node := range queue { // 记得。层次遍历用这个方式来遍历！
			if idx == len(queue)-1 {
				ans = append(ans, node.Val)
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right) // 记得别写错啦，是node
			}
			queue = q // 记得赋值
		}
	}
	return ans
}
