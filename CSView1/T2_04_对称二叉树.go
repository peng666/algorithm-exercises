package CSView1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	return isInvertTrees(root.Left, root.Right)
}
func isInvertTrees(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true // 记清楚这些结束条件！！
	}
	if (root1 == nil && root2 != nil) || (root1 != nil && root2 == nil) {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	return isInvertTrees(root1.Left, root2.Right) && isInvertTrees(root1.Right, root2.Left)
	记得两边都是true才是true
}
