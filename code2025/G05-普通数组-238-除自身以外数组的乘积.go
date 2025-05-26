package main

import "fmt"

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1
	tmp := 1 // 记录右边的乘积
	for i := 1; i < len(res); i++ {
		res[i] = res[i-1] * nums[i-1] // 先计算左边的乘积
	}

	for i := len(res) - 2; i >= 0; i-- {
		tmp *= nums[i+1]
		res[i] *= tmp
	}
	return res
}

func main() {
	nums := []int{1, 2, 3, 4}
	res := productExceptSelf(nums)
	fmt.Println(res)
}
