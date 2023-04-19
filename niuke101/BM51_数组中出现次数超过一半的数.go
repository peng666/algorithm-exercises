package main

import "fmt"

/**
 *
 * @param numbers int整型一维数组
 * @return int整型
 */
func MoreThanHalfNum_Solution(numbers []int) int {
	// write code here
	s := len(numbers)
	if s <= 0 {
		return 0
	}
	//halfCount := s/2 + 1
	//m := make(map[int]int)
	//for i := 0; i < s; i++ {
	//	if m[numbers[i]] == 0 {
	//		m[numbers[i]] = 1
	//	} else {
	//		m[numbers[i]]++
	//	}
	//	if m[numbers[i]] >= halfCount {
	//		return numbers[i]
	//	}
	//}
	//return 0

	// 以下是最优解，候选法（不需要内存）
	cond := -1
	cnt := 0
	for i := 0; i < s; i++ {
		if numbers[i] == cond {
			cnt++
		} else {
			if cnt == 0 {
				cond = numbers[i]
				cnt++
			} else {
				cnt--
			}
		}
	}
	return cond
}

func main() {
	nums := []int{3}
	ret := MoreThanHalfNum_Solution(nums)
	fmt.Println(ret)
}
