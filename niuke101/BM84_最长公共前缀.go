//package main
//
//import "fmt"
//
///**
// *
// * @param strs string字符串一维数组
// * @return string字符串
// */
//func longestCommonPrefix(strs []string) string {
//	// write code here
//	l := len(strs)
//	if l <= 0 {
//		return ""
//	}
//	if l == 1 {
//		return strs[0]
//	}
//
//	maxLen := len(strs[0])
//	for i := 1; i < l; i++ {
//		for j := 0; j < maxLen; j++ {
//			if j >= len(strs[i]) {
//				maxLen = j
//				break
//			}
//
//			if strs[0][j] == strs[i][j] {
//				continue
//			} else {
//				maxLen = j
//			}
//		}
//	}
//	return strs[0][0:maxLen]
//}
//
//func main() {
//	fmt.Println("123")
//	strs := []string{"abc"}
//	result := longestCommonPrefix(strs)
//	fmt.Println(result)
//}
