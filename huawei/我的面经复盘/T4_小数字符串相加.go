package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
题干描述将两个小数格式的字符串相加，输出两个小数相加的结果（不能使用String转数字的库函数、大数处理函数）

输入：a=‘456.456’b=‘2354.932446’
输出：2811.388446
*/

func add(s1, s2 string) string {
	if s1 == "" {
		return s2
	}
	if s2 == "" {
		return s1
	}

	num1 := strings.Split(s1, ".")
	num2 := strings.Split(s2, ".")
	num1Z := num1[0] //Z整数部分，X小数部分
	num1X := ""
	num2Z := num2[0]
	num2X := ""
	if len(num1) == 2 {
		num1X = num1[1]
	}
	if len(num2) == 2 {
		num2X = num2[1]
	}

	//小数部分的最大长度
	lenX := max(len(num1X), len(num2X))

	// 长度不同就使之长度相等
	if len(num1X) != len(num2X) {
		if len(num1X) > len(num2X) {
			c := len(num1X) - len(num2X)
			for c > 0 {
				num2X = num2X + "0"
				c--
			}
		}
		if len(num1X) < len(num2X) {
			c := len(num2X) - len(num1X)
			for c > 0 {
				num1X = num1X + "0"
				c--
			}
		}
	}

	numZStr := addTwoNum(num1Z, num2Z)
	numXStr := addTwoNum(num1X, num2X)

	// 小数部分转字符串
	if len(numXStr) > lenX {
		numZStr = addTwoNum(numZStr, numXStr[0:1])
		numXStr = numXStr[1:]
	} else if len(numXStr) < lenX {
		c := lenX - len(numXStr)
		for c > 0 {
			numXStr = "0" + numXStr
			c--
		}
	}

	//去掉小数末尾的0
	idx := len(numXStr)
	for idx > 0 && numXStr[idx-1] == '0' {
		idx--
	}
	numXStr = numXStr[:idx]

	// 拼接整数部分和小数部分
	if len(numXStr) <= 0 {
		return numZStr
	}
	return numZStr + "." + numXStr
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func addTwoNum(s1, s2 string) string {
	res := ""
	sum := 0
	i := len(s1) - 1
	j := len(s2) - 1

	for i >= 0 || j >= 0 {
		num1 := 0
		num2 := 0

		if i >= 0 {
			num1 = int(s1[i] - '0')
			i--
		}
		if j >= 0 {
			num2 = int(s2[j] - '0')
			j--
		}

		sum = sum + num1 + num2
		res = strconv.Itoa(sum%10) + res
		sum = sum / 10
	}
	if sum != 0 {
		res = strconv.Itoa(sum%10) + res
	}
	return res
}
func main() {
	s1 := "11.11"
	s2 := "1.9"
	result := add(s1, s2)
	fmt.Println(result)
}
