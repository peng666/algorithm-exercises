package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param str string字符串 待判断的字符串
 * @return bool布尔型
 */
func judge(str string) bool {
	// write code here

	l := len(str)
	if l <= 1 {
		return true
	}

	left, right := 0, l-1
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	str := "qweerwq"
	fmt.Println(judge(str))
}
