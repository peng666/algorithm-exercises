package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	res := 0
	left := 0
	hashmap := make(map[byte]int)
	for i, _ := range s {
		if idx, ok := hashmap[s[i]]; ok {
			if idx >= left {
				left = idx + 1
			}
		}
		hashmap[s[i]] = i
		res = max(res, i-left+1)
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
	res := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(res)
}
