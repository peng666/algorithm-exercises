package main

import "fmt"

func MergeSort(nums *[]int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	MergeSort(nums, start, mid)
	MergeSort(nums, mid+1, end)
	if (*nums)[mid] > (*nums)[mid+1] {
		MergeNums(nums, start, mid, end)
	}
}

func MergeNums(nums *[]int, start, mid, end int) {
	tmp := []int{} // 注意这个的初始化大小和值！

	i, j := start, mid+1
	for i <= mid && j <= end {
		if (*nums)[i] <= (*nums)[j] {
			tmp = append(tmp, (*nums)[i])
			i++
		} else {
			tmp = append(tmp, (*nums)[j])
			j++
		}
	}
	if i <= mid {
		tmp = append(tmp, (*nums)[i:mid+1]...)
	}
	if j <= end {
		tmp = append(tmp, (*nums)[j:end+1]...)
	}

	for idx, num := range tmp {
		(*nums)[idx+start] = num
	}
}

func main() {
	nums := []int{2, 5, 6, 3, 43, 23, 9, 6}
	MergeSort(&nums, 0, len(nums)-1) // 不应该传这么多参数进去，应该就只给个数组
	fmt.Println(nums)
}
