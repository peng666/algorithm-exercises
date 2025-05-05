package CSView1

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	// 方法1：回溯+哈希。遍历一遍链表，边创建新链表，边用map保存映射关系。然后再遍历链表设置random指针
	// 方法2：迭代+节点拆分。不需要额外的存储空间。遍历3遍。创建新节点；指定random指针；拆分新链表
	cur := head
	for cur != nil {
		tmp := &Node{Val: cur.Val, Next: cur.Next}
		cur.Next = tmp
		cur = cur.Next.Next
	}
	cur = head
	for cur != nil {
		if cur.Random == nil {
			cur.Next.Random = nil
		} else {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	cur = head
	dummy := new(Node)
	pre := dummy
	for cur != nil {
		pre.Next = cur.Next
		cur.Next = cur.Next.Next
		pre = pre.Next
		cur = cur.Next
	}

	return dummy.Next
}
