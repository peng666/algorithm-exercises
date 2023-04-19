package main

import "fmt"

/**
 * max water
 * @param arr int整型一维数组 the array
 * @return long长整型
 */
func maxWater(arr []int) int64 {
	// write code here
	s := len(arr)
	if s <= 1 {
		return 0
	}

	l, r := 0, s-1
	ret := 0
	maxL, maxR := 0, 0
	for l < r {
		if arr[l] > maxL {
			maxL = arr[l]
		}
		if arr[r] > maxR {
			maxR = arr[r]
		}
		if maxL < maxR {
			ret += maxL - arr[l]
			l++
		} else {
			ret += maxR - arr[r]
			r--
		}
	}
	return int64(ret)
}

func main() {
	arr := []int{42}
	fmt.Println(maxWater(arr))
}
