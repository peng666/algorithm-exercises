package main

/*
求一个字符串，分割出子串（不重复）的最多数量
思路：滑动窗口，使用map维护重复记录
*/
import "fmt"

func maxSplit(s string) int {
	n := len(s)
	if n <= 0 {
		return 0
	}
	// 存储已经出现过的字符串
	m := make(map[string]bool)

	// 双指针遍历字符串---滑动窗口
	m[s[0:1]] = true
	maxCount := 1
	i, j := 1, 2
	for j <= n {
		if m[s[i:j]] {
			j++
			continue
		}
		m[s[i:j]] = true
		maxCount++
		i = j
		j++
	}
	return maxCount
}

func main() {
	s := "ababccc"
	res := maxSplit(s)
	fmt.Println(res) //5
}
