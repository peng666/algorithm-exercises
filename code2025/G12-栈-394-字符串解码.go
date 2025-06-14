package main

import "fmt"

func decodeString(s string) string {
	var numStack []int
	var strStack []string
	num := 0
	temp := ""
	for idx, _ := range s {
		if s[idx] >= '0' && s[idx] <= '9' { // 暂存数字
			num = num*10 + int(s[idx]-'0')
		} else if s[idx] == '[' { // 数字、字符串压栈、清零
			numStack = append(numStack, num)
			strStack = append(strStack, temp)
			num = 0
			temp = ""
		} else if s[idx] == ']' { // 数字、字符串出栈，获取结果
			count := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			var str string
			for count > 0 {
				str = str + temp
				count -= 1
			}
			preStr := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			temp = preStr + str
		} else { // 记录临时字符串
			temp = temp + string(s[idx])
		}
	}
	return temp
}

func main() {
	res := decodeString("3[a2[c]]")
	fmt.Print(res)
}
