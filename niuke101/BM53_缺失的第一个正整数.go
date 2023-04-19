package main

import (
	"fmt"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param nums int整型一维数组
 * @return int整型
 */
func minNumberDisappeared(nums []int) int {
	// write code here
	n := len(nums)
	if n <= 0 {
		return 0
	}

	// 可以使用map保存每一个数字出现的次数，然后遍历map就知道是谁了
	// 也可以原地map
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	for i := 0; i < n; i++ {
		tmp := abs(nums[i])
		if tmp <= n {
			nums[tmp-1] = -abs(nums[tmp-1]) // 注意我要它为负数
		}

	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}
func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	nums := []int{4, 5, 6, 8, 9}
	ret := minNumberDisappeared(nums)
	fmt.Println(ret)
}
