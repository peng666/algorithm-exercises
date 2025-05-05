package CSView1

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := new(ListNode) // 要操作源链表却不知道头节点是否要动，可以建一个虚节点~~
	dummy.Next = head

	pre := dummy
	for i := 1; i < left; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	end := cur
	var newHead, next *ListNode
	for i := 0; i < right-left+1; i++ { // 记住要+1，理清楚这些临界条件
		next = cur.Next
		cur.Next = newHead
		newHead = cur
		cur = next
	}
	pre.Next = newHead
	end.Next = next
	return dummy.Next
}
