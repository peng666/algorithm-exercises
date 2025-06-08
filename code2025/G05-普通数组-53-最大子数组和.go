package main

import "fmt"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]
	ans := 0
	for _, num := range nums {
		ans += num
		if ans > res {
			res = ans
		}
		if ans < 0 {
			ans = 0
		}
	}
	return res
}

func main() {
	res := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Println(res)
}
