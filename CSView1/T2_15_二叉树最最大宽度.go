package CSView1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type item struct {
	idx      int
	treeNode *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	queue := []*item{{0, root}} // 记录满二叉树的编号

	for len(queue) > 0 {
		width := queue[len(queue)-1].idx - queue[0].idx + 1
		if width > ans {
			ans = width
		}

		q := []*item{} // 跟层次遍历一样，就是多存个树节点的编号idx
		for _, node := range queue {
			if node.treeNode.Left != nil {
				q = append(q, &item{idx: node.idx * 2, treeNode: node.treeNode.Left})
			}
			if node.treeNode.Right != nil {
				q = append(q, &item{idx: node.idx*2 + 1, treeNode: node.treeNode.Right})
			}
		}
		queue = q
	}
	return ans
}
