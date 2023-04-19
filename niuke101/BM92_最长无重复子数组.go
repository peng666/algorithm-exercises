package main

import "fmt"

/**
 *
 * @param arr int整型一维数组 the array
 * @return int整型
 */
func maxLength(arr []int) int {
	// write code here
	s := len(arr)
	if s == 0 {
		return 0
	}
	if s == 1 {
		return 1
	}
	l, r := 0, 0
	m := make(map[int]int)
	result := 0
	tmp := 0
	for r < s {
		if m[arr[r]] == 0 {
			m[arr[r]] = 1
			r++
			tmp++
			if tmp > result {
				result = tmp
			}
		} else {
			m[arr[l]]--
			l++
			tmp--
		}
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 1, 2, 3, 2, 2}
	ret := maxLength(nums)
	fmt.Println(ret)
}
