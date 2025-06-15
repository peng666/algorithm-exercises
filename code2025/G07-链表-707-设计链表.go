package main

type Node struct {
	Val       int
	Pre, Next *Node
}

type MyLinkedList struct {
	Head, Tail *Node
	size       int
}

func Constructor() MyLinkedList {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Pre = head

	return MyLinkedList{
		Head: head,
		Tail: tail,
		size: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	cur := this.Head.Next
	i := 0
	for cur != this.Tail {
		if i == index {
			return cur.Val
		}
		i += 1
		cur = cur.Next
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{Val: val, Pre: this.Head, Next: this.Head.Next}
	this.Head.Next.Pre = node
	this.Head.Next = node
	this.size += 1
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &Node{Val: val, Pre: this.Tail.Pre, Next: this.Tail}
	this.Tail.Pre.Next = node
	this.Tail.Pre = node
	this.size += 1
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.size {
		return
	}
	cur := this.Head.Next
	i := 0
	for cur != nil {
		if i == index {
			node := &Node{Val: val, Pre: cur.Pre, Next: cur}
			cur.Pre.Next = node
			cur.Pre = node
			this.size += 1
		}
		i += 1
		cur = cur.Next
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	cur := this.Head.Next
	i := 0
	for cur != this.Tail {
		if i == index {
			cur.Pre.Next = cur.Next
			cur.Next.Pre = cur.Pre
			this.size -= 1
		}
		i += 1
		cur = cur.Next
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
