package main

import "fmt"

type Node struct {
	Key   int
	Value int
	Next  *Node
	Pre   *Node
}

type Solution struct {
	hash       map[int]*Node
	capacity   int
	size       int
	head, tail *Node
}

func Constructor(capacity int) Solution {
	return Solution{
		hash:     make(map[int]*Node, capacity),
		capacity: capacity,
	}
}

func (s *Solution) get(key int) int {
	node, ok := s.hash[key]
	if !ok {
		return -1
	}
	if node != s.tail {
		if node.Pre != nil {
			node.Pre.Next = node.Next
		} else {
			s.head = node.Next
		}
		node.Next.Pre = node.Pre

		s.tail.Next = node
		node.Pre = s.tail
		s.tail = node
	}
	return node.Value
}
func (s *Solution) set(key int, value int) {
	v := s.get(key)
	if v == -1 {
		if s.size == s.capacity {
			delete(s.hash, s.head.Key)
			s.head = s.head.Next
			s.head.Pre = nil
		} else {
			s.size++
		}
		newNode := &Node{
			Key:   key,
			Value: value,
		}
		s.hash[key] = newNode
		if s.tail == nil {
			s.tail = newNode
			s.head = newNode
		} else {
			s.tail.Next = newNode
			s.tail = newNode
		}
	} else {
		s.tail.Value = value
	}
	return
}

func main() {
	solution := Constructor(1)
	solution.set(1, 2)
	solution.set(1, 3)
	solution.set(2, 4)
	solution.set(3, 9)
	fmt.Println(solution.get(1))
	fmt.Println(solution.get(2))
	solution.set(2, 12)
	fmt.Println(solution.get(2))
}
