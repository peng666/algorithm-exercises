package main

import "fmt"

func trap(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := 0
	left, right := 0, len(nums)-1
	leftMax, rightMax := nums[left], nums[right]
	for left < right { // 这里注意不能等于。因为里面是先移动指针，后计算的
		if leftMax < rightMax {
			left += 1
			cur := nums[left]
			if cur > leftMax {
				leftMax = cur
			} else {
				res += leftMax - cur
			}
		} else {
			right -= 1
			cur := nums[right]
			if cur > rightMax {
				rightMax = cur
			} else {
				res += rightMax - cur
			}
		}
	}

	return res
}

func main() {
	res := trap([]int{4, 2, 0, 3, 2, 5})
	fmt.Print(res)
}
