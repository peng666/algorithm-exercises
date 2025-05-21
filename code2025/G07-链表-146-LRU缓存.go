package main

import "fmt"

type Node struct {
	Key       int
	Val       int
	Pre, Next *Node
}

type LRUCache struct {
	Capacity   int
	hashmap    map[int]*Node
	Head, Tail *Node
}

func NewLRUCache(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Pre = head

	return LRUCache{
		Capacity: capacity,
		hashmap:  make(map[int]*Node, capacity),
		Head:     head,
		Tail:     tail,
	}
}

func (l *LRUCache) Get(key int) int {
	if node, ok := l.hashmap[key]; ok {
		l.moveToHead(node)
		return node.Val
	}
	return -1
}

func (l *LRUCache) Put(key, val int) {
	if l.Get(key) != -1 {
		l.hashmap[key].Val = val
		return
	}

	node := &Node{Key: key, Val: val}
	if len(l.hashmap) == l.Capacity {
		l.removeTail()
	}
	l.hashmap[key] = node
	l.addToHead(node)
}

func (l *LRUCache) moveToHead(node *Node) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
	l.addToHead(node)
}

func (l *LRUCache) removeTail() {
	node := l.Tail.Pre
	delete(l.hashmap, node.Key)
	l.Tail.Pre = node.Pre
	node.Pre.Next = l.Tail
}

func (l *LRUCache) addToHead(node *Node) {
	l.Head.Next.Pre = node
	node.Next = l.Head.Next
	node.Pre = l.Head
	l.Head.Next = node
}

func main() {
	cache := NewLRUCache(3)
	fmt.Println(cache.Get(1))

	cache.Put(1, 10)
	cache.Put(2, 20)
	cache.Put(3, 30)

	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(2))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))

	cache.Put(4, 40)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(2))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
}
