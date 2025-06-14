package main

import "fmt"

func longestValidParentheses(s string) int {
	res := 0
	dp := make([]int, len(s)) // 以i结尾的最大长度
	for i := 0; i < len(s); i += 1 {
		length := 0
		if i > 0 && s[i] == ')' {
			if s[i-1] == '(' {
				if i > 1 {
					length = dp[i-2] + 2
				} else {
					length = 2
				}
			} else if s[i-1] == ')' && i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
				length = dp[i-1] + 2
				if i-dp[i-1]-2 >= 0 {
					length += dp[i-dp[i-1]-2]
				}
			}
		}
		dp[i] = length
		if length > res {
			res = length
		}
	}
	return res
}

func main() {
	re := longestValidParentheses("()(())")
	fmt.Print(re)
}
