package niuke101b

/*
	先分割单词，然后从后往前遍历（倒序遍历），一个个改大小写（一个个字母添加回去，所以记得加上空格）
*/

import "strings"

func trans(s string, n int) string {
	words := strings.Split(s, " ") // 先分割出单词
	res := make([]byte, n)
	count := 0
	for i := len(words) - 1; i >= 0; i-- {
		str := words[i]
		for j := 0; j < len(str); j++ {
			if str[j] >= 'a' && str[j] <= 'z' {
				res[count] = str[j] + 'A' - 'a'
			} else {
				res[count] = str[j] - 'A' + 'a'
			}
			count++
		}
		if count < n {
			res[count] = ' '
		}
		count++
	}
	return string(res)
}
