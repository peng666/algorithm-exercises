package main

/*
思路：
	一个capacity，双链表（head，tail），哈希表map[int]*Node。
还要定义Node（key，val，前后指针pre，next）。
3个函数：construction记住head\tail空节点，get很简单，put记住删除map的
辅助：moveToHead,deleteTail,addToHead
*/

type LRUCache struct {
	capacity   int
	hashTable  map[int]*Node
	head, tail *Node
}

type Node struct {
	key, val  int
	pre, next *Node
}

func Constructor(cap int) LRUCache {
	head, tail := &Node{}, &Node{} // 空节点，方便
	head.next = tail
	tail.pre = head
	return LRUCache{
		capacity:  cap,
		hashTable: make(map[int]*Node, cap),
		head:      head,
		tail:      tail,
	}
}

func (this *LRUCache) moveToHead(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
	this.addToHead(node)
}

func (this *LRUCache) deleteTail() {
	k := this.tail.pre.key
	this.tail.pre.pre.next = this.tail
	this.tail.pre = this.tail.pre.pre
	delete(this.hashTable, k)
}

func (this *LRUCache) addToHead(node *Node) {
	node.pre = this.head
	node.next = this.head.next
	this.head.next.pre = node
	this.head.next = node
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.hashTable[key]; ok {
		this.moveToHead(node)
		return node.val
	}
	return -1
}
func (this *LRUCache) Put(key, val int) {
	if tmp := this.Get(key); tmp != -1 {
		this.head.next.val = val
		return
	}
	if this.capacity == len(this.hashTable) {
		this.deleteTail()
	}
	node := &Node{
		key: key,
		val: val,
	}
	this.addToHead(node)
	this.hashTable[key] = node
}
