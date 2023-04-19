package main

import "fmt"

// 快速排序
// 注意：1，数组传引用；2，记清楚左闭右开的两指针；3，快排以左边为基准，就先移动右指针
func quickSort(nums []int) []int {
	quick(&nums, 0, len(nums))
	return nums
}

func quick(nums *[]int, left, right int) {
	if left >= right {
		return
	}
	pivot := (*nums)[left]
	i, j := left, right-1
	for i < j {
		for (*nums)[j] >= pivot && i < j {
			j--
		}
		for (*nums)[i] <= pivot && i < j {
			i++
		}
		(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
	}
	(*nums)[i], (*nums)[left] = (*nums)[left], (*nums)[i]
	quick(nums, left, i)
	quick(nums, i+1, right)
	return
}

func main() {
	nums := []int{2, 3, 5, 1, 2, 6}
	quickSort(nums)
	fmt.Println(nums)
}
