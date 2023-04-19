package main

import (
	"fmt"
	"sort"
)

/**
 *
 * @param num int整型一维数组
 * @return int整型二维数组
 */
func threeSum(num []int) [][]int {
	// write code here
	n := len(num)
	if n <= 2 {
		return nil
	}

	res := [][]int{}
	//先排序，因为答案需要有序
	sort.Ints(num)
	for i := 0; i < n-2; i++ {
		if num[i] > 0 {
			break // 因为已经是有序，提前结束
		}

		if i > 0 && num[i] == num[i-1] { // 这里不太对吧？！[-1,-1,2]。这是跟前一个比较！注意
			continue // 去重
		}

		for l, r := i+1, n-1; l < r; {
			tmp := num[i] + num[l] + num[r]
			if tmp == 0 {
				res = append(res, []int{num[i], num[l], num[r]})

				for l < r && num[l] == num[l+1] {
					l++ // 去重
				}
				l++
				r--
			} else if tmp < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}

func main() {
	nums := []int{0, 0, 0, 0}
	res := threeSum(nums)
	fmt.Println(res)
}
