package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 0
	fmt.Scan(&n)

	m := map[int]bool{}
	numArr := []int{}

	for i := 0; i < n; i++ {
		num := 0
		fmt.Scan(&num)
		if m[num] == true {
			continue
		}
		m[num] = true
		numArr = append(numArr, num)
	}
	sort.Ints(numArr)
	for _, num := range numArr {
		fmt.Println(num)
	}
}
