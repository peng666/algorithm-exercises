package main

import "fmt"

func singleNumber(nums []int) []int {
	temp := 0
	for _, num := range nums {
		temp ^= num
	}

	// 获取 000001000000。用来将两个数分组，1的位置都是两数末尾的不同之处，下面&与运算就分组了
	temp = temp & -temp // 负数是按位取反再加1

	// 分组求出
	res1, res2 := 0, 0
	for _, num := range nums {
		if num&temp == 0 {
			res1 ^= num
		} else {
			res2 ^= num
		}
	}
	return []int{res1, res2}
}

func main() {
	res := singleNumber([]int{2, 3, 4, 2, 6, 7, 7, 4})
	fmt.Println(res)
}
