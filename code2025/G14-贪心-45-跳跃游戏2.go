package main

func jump(nums []int) int {
	curDistance := 0
	maxDistance := 0
	ans := 0
	for i, num := range nums {
		if i > curDistance {
			curDistance = maxDistance
			ans += 1
		}
		if i <= curDistance {
			if i+num > maxDistance {
				maxDistance = i + num
			}
		}
	}
	return ans
}
