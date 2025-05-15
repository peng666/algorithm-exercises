package main

import "fmt"

func singleNumber(nums []int) int {
	res := 0
	for _, num := range nums {
		res = res ^ num
	}
	return res
}

func main() {
	fmt.Println(singleNumber([]int{3, 6, 4, 3, 4, 8, 8}))
}
