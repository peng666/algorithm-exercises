package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param array int整型一维数组
 * @return int整型一维数组
 */
func FindNumsAppearOnce(array []int) []int {
	// write code here
	s := len(array)
	if s <= 1 {
		return []int{}
	}
	// 可以用map
	// 以下利用异或运算
	temp := 0
	for i := 0; i < s; i++ {
		temp ^= array[i]
	}
	k := 1
	for k&temp == 0 {
		k <<= 1 // 找出两数的某个不同的位
	}
	res1, res2 := 0, 0
	for i := 0; i < s; i++ {
		if k&array[i] == 0 {
			res1 ^= array[i]
		} else {
			res2 ^= array[i]
		}
	}
	if res1 < res2 {
		return []int{res1, res2}
	} else {
		return []int{res2, res1}
	}
	return []int{}
}

func main() {
	nums := []int{1, 2, 3, 3, 2, 9}
	res := FindNumsAppearOnce(nums)
	if len(res) > 1 {
		fmt.Println(res[0], res[1])
	}

}
