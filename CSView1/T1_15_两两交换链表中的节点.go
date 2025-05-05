package CSView1

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Val: 0, Next: head} // 注意这里，指向头节点
	cur := dummy
	for cur != nil && cur.Next != nil && cur.Next.Next != nil {
		node1 := cur.Next
		node2 := cur.Next.Next
		node1.Next = node2.Next
		node2.Next = node1
		cur.Next = node2
		cur = cur.Next.Next
	}
	return dummy.Next
}
