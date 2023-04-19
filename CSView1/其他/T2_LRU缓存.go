//package main
//
//// 一个Node（kv，前后指针）
//// 一个LRUCache(一个hash哈希表，一个双向链表，两个属性大小和容量)
//
//type Node struct {
//	Key       int
//	Val       int
//	Pre, Next *Node
//}
//
//type LRUCache struct {
//	size       int
//	capacity   int
//	hash       map[int]*Node
//	head, tail *Node
//}
//
//func Constructor(cap int) LRUCache {
//	return LRUCache{
//		size:     0,
//		capacity: cap,
//		hash:     make(map[int]*Node, cap),
//	}
//}
//
//func (l *LRUCache) Get(key int) int {
//	node, ok := l.hash[key]
//	if !ok {
//		return -1
//	}
//
//	if node != l.tail {
//		if node.Pre != nil {
//			node.Pre.Next = node.Next
//		} else {
//			l.head = node.Next
//		}
//		node.Next.Pre = node.Pre
//		l.tail.Next = node
//		node.Pre = l.tail
//		l.tail = node
//	}
//	return node.Val
//}
//
//func (l *LRUCache) Put(key, val int) {
//	v := l.Get(key)
//	if v == -1 {
//		if l.size == l.capacity {
//			delete(l.hash, l.head.Key)
//			l.head = l.head.Next
//			l.head.Pre = nil
//		} else {
//			l.size++
//		}
//		node := &Node{
//			Key:  key,
//			Val:  val,
//			Pre:  l.tail,
//			Next: nil,
//		}
//		l.hash[key] = node
//		if l.tail == nil {
//			l.tail = node
//			l.head = node
//		} else {
//			l.tail.Next = node
//			l.tail = node
//		}
//	} else {
//		l.tail.Val = val
//	}
//}
