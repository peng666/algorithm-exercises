package main

import (
	"sort"
	"strconv"
)

// 给定一个数组，拼接元素输出最大数，不能拆分元素里的数
// 思路：有点像冒泡排序
func solve(nums []int) string {
	tmp := []string{}

	for _, v := range nums {
		tmp = append(tmp, strconv.Itoa(v))
	}
	// 给元素们排序，把能组成最大值的放前面
	sort.Slice(tmp, func(i, j int) bool {
		num1, _ := strconv.Atoi(tmp[i] + tmp[j])
		num2, _ := strconv.Atoi(tmp[j] + tmp[i])
		return num1 > num2
	})

	res := ""
	for _, v := range tmp {
		res += v
	}
	if res[0] == '0' {
		return "0"
	}
	return res

}
