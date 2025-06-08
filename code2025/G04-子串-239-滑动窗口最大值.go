package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	var queue []int
	for idx, num := range nums {
		for len(queue) > 0 && queue[0] <= idx-k {
			queue = queue[1:]
		}

		for len(queue) > 0 && num > nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, idx)

		if idx >= k-1 {
			res = append(res, nums[queue[0]])
		}
	}
	return res
}

func main() {
	res := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	fmt.Println(res)

}
