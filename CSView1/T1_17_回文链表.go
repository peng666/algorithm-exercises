package CSView1

type ListNode struct {
	Val  int
	Next *ListNode
}

等于两到小题组合：快慢指针找中点；反转链表
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	if head.Next == nil {
		return true
	}

	// 快慢指针。先找中点
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	head2 := reverse(slow)
	for head2 != nil {
		if head.Val != head2.Val {
			return false
		}
		head2 = head2.Next // 遍历链表记得要移动指针呀！你老是忘记
		head = head.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var newHead *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = newHead
		newHead = head
		head = tmp
	}
	return newHead
}
