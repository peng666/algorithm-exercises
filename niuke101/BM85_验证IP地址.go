//package main
//
//import (
//	"fmt"
//	"strconv"
//	"strings"
//)
//
///**
// * 验证IP地址
// * @param IP string字符串 一个IP地址字符串
// * @return string字符串
// */
//func solve(IP string) string {
//	// write code here
//	if IP == "" {
//		return "Neither"
//	}
//
//	strArr := strings.Split(IP, ".")
//	l := len(strArr)
//
//	if l == 4 {
//		for i := 0; i < 4; i++ {
//			if len(strArr[i]) > 3 || len(strArr[i]) <= 0 {
//				return "Neither"
//			}
//			if strArr[i][0] == '0' {
//				return "Neither"
//			}
//			for j := 0; j < len(strArr[i]); j++ {
//				if strArr[i][j] > '9' || strArr[i][j] < '0' {
//					return "Neither"
//				}
//			}
//			num, _ := strconv.Atoi(strArr[i])
//			if num < 0 || num > 255 {
//				return "Neither"
//			}
//		}
//		return "IPv4"
//	}
//
//	strArr = strings.Split(IP, ":")
//	l = len(strArr)
//
//	if l == 8 {
//		for i := 0; i < 8; i++ {
//			if len(strArr[i]) > 4 || len(strArr[i]) <= 0 {
//				return "Neither"
//			}
//			for j := 0; j < len(strArr[i]); j++ {
//				if !((strArr[i][j] >= '0' && strArr[i][j] <= '9') || (strArr[i][j] >= 'a' && strArr[i][j] <= 'f') || (strArr[i][j] >= 'A' && strArr[i][j] <= 'F')) {
//					return "Neither"
//				}
//			}
//		}
//		return "IPv6"
//	}
//	return "Neither"
//}
//
//func main() {
//	fmt.Println("123")
//	str := "256.256.256.256"
//	result := solve(str)
//	fmt.Println(result)
//}
