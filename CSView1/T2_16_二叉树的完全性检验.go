//package CSView1
//
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//type item struct {
//	idx int
//	*TreeNode
//}
//
//func isCompleteTree(root *TreeNode) bool {
//	if root == nil {
//		return true
//	}
//	cnt := 1 // 记录节点序号
//	queue := []*item{{1, root}}
//	for len(queue) > 0 {
//		q := []*item{}
//		for _, node := range queue {
//			if node.idx != cnt { // 记清楚这个idx是item里的idx！不是本次迭代的idx
//				return false
//			}
//			if node.Left != nil {
//				q = append(q, &item{node.idx * 2, node.Left})
//			}
//			if node.Right != nil {
//				q = append(q, &item{node.idx*2 + 1, node.Right})
//			}
//			cnt++
//		}
//		queue = q
//	}
//	return true
//}
