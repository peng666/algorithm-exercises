package main

import (
	"fmt"
)

/**
 *
 * @param numbers int整型一维数组
 * @param target int整型
 * @return int整型一维数组
 */
func twoSum(numbers []int, target int) []int {
	// write code here
	s := len(numbers)
	if s <= 1 {
		return []int{}
	}

	m := make(map[int]int)
	for i := 0; i < s; i++ {
		if m[target-numbers[i]] == 0 {
			m[numbers[i]] = i + 1
		} else {
			return []int{m[target-numbers[i]], i + 1}
		}
	}
	return []int{}
}

func main() {
	nums := []int{3, 2, 4}
	ret := twoSum(nums, 6)
	fmt.Println(ret)
}
