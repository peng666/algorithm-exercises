package main

import "fmt"

/**
 *
 * @param A int整型一维数组
 * @param B int整型一维数组
 * @return void
 */
func merge(A []int, m int, B []int, n int) {
	// write code here

	l := m + n
	p1 := m - 1
	p2 := n - 1
	for i := l - 1; i >= 0 && p1 >= 0 && p2 >= 0; i-- {
		if A[p1] > B[p2] {
			A[i] = A[p1]
			p1--
		} else {
			A[i] = B[p2]
			p2--
		}
	}
	for p2 >= 0 {
		A[p2] = B[p2]
		p2--
	}
}

func main() {
	A := []int{1, 3, 4, 6, 8, 0, 0, 0, 0, 0, 0}
	B := []int{2, 4, 5, 6, 7, 12}
	merge(A, 5, B, 6)
	fmt.Println(A)
}
