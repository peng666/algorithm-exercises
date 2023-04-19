package main

type LRUCache struct {
	capacity   int
	m          map[int]*Node
	head, tail *Node
}

type Node struct {
	Key       int
	Val       int
	Pre, Next *Node
}

func Constructor(capacity int) LRUCache {
	head, tail := &Node{}, &Node{}
	head.Next = tail
	tail.Pre = head
	return LRUCache{
		capacity: capacity,
		m:        make(map[int]*Node, capacity),
		head:     head, // 弄个空节点，可以不用去判断空
		tail:     tail,
	}
}

func (this *LRUCache) deleteNode(node *Node) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (this *LRUCache) addToHead(node *Node) {
	node.Pre = this.head
	node.Next = this.head.Next
	this.head.Next.Pre = node
	this.head.Next = node
}

func (this *LRUCache) moveToHead(node *Node) {
	this.deleteNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() int {
	node := this.tail.Pre
	this.deleteNode(node)
	return node.Key
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.m[key]; ok {
		this.moveToHead(v)
		return v.Val
	}
	return -1
}

func (this *LRUCache) Put(key, val int) {
	if v, ok := this.m[key]; ok {
		v.Val = val
		this.moveToHead(v)
		return
	}
	if this.capacity == len(this.m) {
		delKey := this.removeTail()
		delete(this.m, delKey) // 记住要删除hash里的
	}
	newNode := &Node{
		Key: key,
		Val: val,
	}
	this.addToHead(newNode)
	this.m[key] = newNode
}
