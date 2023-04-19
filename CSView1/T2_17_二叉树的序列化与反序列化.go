//package main
//
//import (
//	"fmt"
//	"strconv"
//	"strings"
//)
//
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//type Codec struct {
//}
//
//func Constructor() Codec {
//	return *new(Codec)
//}
//
//func (c *Codec) serialize(root *TreeNode) string {
//	if root == nil {
//		return "x"
//	}
//	// 记得加上分隔符
//	return strconv.Itoa(root.Val) + "," + c.serialize(root.Left) + "," + c.serialize(root.Right)
//}
//
//func (c *Codec) deserialize(data string) *TreeNode {
//	if data == "" {
//		return nil
//	}
//	lists := strings.Split(data, ",")
//	return c.buildTree(&lists)
//}
//
//func (c *Codec) buildTree(list *[]string) *TreeNode {
//	if len(*list) == 0 {
//		return nil
//	}
//	rootVal := (*list)[0] // 取节点
//	*list = (*list)[1:]   // 记得马上除掉当前节点，因为如果是x就要返回了
//	if rootVal == "x" {
//		return nil
//	}
//	val, _ := strconv.Atoi(rootVal)
//	node := &TreeNode{Val: val}
//	node.Left = c.buildTree(list) // 先不断地递归左子树，剩余的才到右子树（list会动态变，所以要传地址，不能传值）
//	node.Right = c.buildTree(list)
//	return node
//}
//func main() {
//
//	node1 := &TreeNode{Val: 1}
//	node2 := &TreeNode{Val: 2}
//	node3 := &TreeNode{Val: 3}
//	node4 := &TreeNode{Val: 4}
//	node5 := &TreeNode{Val: 5}
//	node1.Left = node2
//	node1.Right = node3
//	node3.Left = node4
//	node3.Right = node5
//
//	ser := Constructor()
//	str := ser.serialize(node1)
//	fmt.Println(str)
//	ser.deserialize(str)
//
//}
