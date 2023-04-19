//package CSView1
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func addTwoNumbers(l1, l2 *ListNode) *ListNode {
//	if l1 == nil {
//		return l2
//	}
//	if l2 == nil {
//		return l1
//	}
//
//	dummy := new(ListNode)
//	cur := dummy
//	num1, num2, tmp := 0, 0, 0
//	for l1 != nil || l2 != nil { // 直接在一个循环里搞定（用或运算）
//		if l1 == nil { // 刚开始就做一些特殊判断
//			num1 = 0
//		} else {
//			num1 = l1.Val
//			l1 = l1.Next
//		}
//		if l2 == nil {
//			num2 = 0
//		} else {
//			num2 = l2.Val
//			l2 = l2.Next
//		}
//		tmp = tmp + num1 + num2
//		node := &ListNode{Val: tmp % 10}
//		tmp = tmp / 10
//		cur.Next = node
//		cur = cur.Next
//	}
//
//	if tmp != 0 {
//		node := &ListNode{Val: tmp}
//		cur.Next = node
//	}
//	return dummy.Next
//}
