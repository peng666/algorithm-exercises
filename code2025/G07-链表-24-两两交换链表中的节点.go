package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre, cur := dummy, dummy.Next
	count := 0
	for head != nil {
		head = head.Next // 指到下一段的开头
		count += 1
		if count == 2 {
			for cur.Next != head {
				tmp := cur.Next // pre固定指开头，cur固定指末尾。tmp指向cur.Next绑在pre.Next处
				cur.Next = tmp.Next
				tmp.Next = pre.Next
				pre.Next = tmp
			}
			pre = cur
			cur = cur.Next
			count = 0
		}
	}
	return dummy.Next
}
