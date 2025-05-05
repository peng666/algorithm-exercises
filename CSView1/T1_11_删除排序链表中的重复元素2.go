package CSView1

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := new(ListNode)
	dummy.Next = head
	pre := dummy
	cur := dummy.Next
	flag := false // 用来标记cur指向的是否为重复
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
			flag = true
			continue
		}
		if flag {
			pre.Next = cur.Next
			cur = pre.Next // 注意：pre、cur两指针记得同时移动，不要忘了其一
			flag = false
		} else {
			pre = pre.Next
			cur = cur.Next
		}
	}
	if flag {
		pre.Next = cur.Next // 注意：最后的重复节点需要再次去校验flag
	}
	return dummy.Next // 注意：dummy记得用来return
}
