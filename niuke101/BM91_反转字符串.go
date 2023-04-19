//package main
//
//import "fmt"
//
///**
// * 反转字符串
// * @param str string字符串
// * @return string字符串
// */
//func solve(str string) string {
//	// write code here
//	if str == "" {
//		return ""
//	}
//
//	tmp := []byte(str)
//	l, r := 0, len(str)-1
//	for l < r {
//		tmp[l] = str[r]
//		tmp[r] = str[l]
//		l++
//		r--
//	}
//	return string(tmp)
//}
//
//func main() {
//	str := "qwertyty"
//	fmt.Println(solve(str))
//}
