package main

import "fmt"

func longestConsecutive(nums []int) int {
	res := 0
	hashmap := make(map[int]bool, len(nums))
	for _, num := range nums {
		hashmap[num] = true
	}

	for num := range hashmap { // 这里要遍历hash表，避免nums有大量重复的首元素，导致求很多遍
		if flag, ok := hashmap[num-1]; ok && flag {
			continue
		}
		cnt := 0
		for hashmap[num+cnt] {
			cnt += 1
		}
		res = max(res, cnt)
	}
	return res
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func main() {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	res := longestConsecutive(nums)
	fmt.Println(res)
}
