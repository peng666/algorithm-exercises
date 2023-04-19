//package main
//
//import (
//	"fmt"
//	"strconv"
//)
//
///**
// * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
// * 计算两个数之和
// * @param s string字符串 表示第一个整数
// * @param t string字符串 表示第二个整数
// * @return string字符串
// */
//func solve(s string, t string) string {
//	// write code here
//
//	tmp, n, m := 0, len(s)-1, len(t)-1
//	result := ""
//	for n >= 0 || m >= 0 || tmp != 0 {
//		i, j := 0, 0
//		if n >= 0 {
//			i = int(s[n] - '0')
//			n--
//		}
//		if m >= 0 {
//			j = int(t[m] - '0')
//			m--
//		}
//		tmp += i + j
//		result = strconv.Itoa(tmp%10) + result
//		tmp = tmp / 10
//	}
//	return result
//}
//
//func main() {
//	s1 := "700077"
//	s2 := "1654"
//	s3 := solve(s1, s2)
//	fmt.Println(s3)
//}
