package main

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) string() string {
	str := strings.Builder{}
	for l != nil {
		str.WriteString(fmt.Sprintf("%d\t", l.Val))
		l = l.Next
	}
	return str.String()
}

func reverseBetween(head *ListNode, m, n int) *ListNode {
	if head == nil || head.Next == nil || m == n {
		return head
	}

	dummy := &ListNode{0, head} // 创建一个临时头节点
	prev := dummy               // 前置节点

	for i := 0; i < m-1; i++ {
		prev = prev.Next // prev之前的不用反转
	}

	pre, cur := &ListNode{}, prev.Next // 从cur开始要反转，pre是新的链表的头
	for i := 0; i < n-m+1; i++ {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	prev.Next.Next = cur
	pre.Next = pre
	return dummy.Next
}

func main() {
	node5 := &ListNode{Val: 5, Next: nil}
	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}

	fmt.Printf("%+v\n", node1.string())
	res := reverseBetween(node1, 2, 4)
	fmt.Printf("%+v", res.string())
}
