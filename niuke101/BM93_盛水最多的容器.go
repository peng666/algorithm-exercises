package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param height int整型一维数组
 * @return int整型
 */
func maxArea(height []int) int {
	// write code here
	s := len(height)
	if s <= 1 {
		return 0
	}

	l, r := 0, s-1
	ret := 0
	for l < r {
		tmp := 0
		if height[l] > height[r] {
			tmp = (r - l) * height[r]
			r-- // 贪心法：每次移动短木板那边
		} else {
			tmp = (r - l) * height[l]
			l++
		}
		if tmp > ret {
			ret = tmp
		}
	}
	return ret
}

func main() {
	h := []int{5, 4, 3, 2, 1, 5}
	fmt.Println(maxArea(h))
}
