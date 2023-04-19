//package CSView1
//
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//func buildTree(preorder, inorder []int) *TreeNode {
//	if len(preorder) == 0 {
//		return nil
//	}
//
//	val := preorder[0]
//	node := &TreeNode{Val: val}
//
//	leftNum := 0
//	for idx, v := range inorder {
//		if v == val {
//			leftNum = idx
//		}
//	}
//
//	node.Left = buildTree(preorder[1:1+leftNum], inorder[0:leftNum])
//	node.Right = buildTree(preorder[1+leftNum:], inorder[leftNum+1:])	// 记得中序遍历序列中间有根节点，要+1
//
//	return node
//}
