package main

func longestPalindrome(s string) string {
	// 1.初始化。定义状态
	dp := make([][]bool, len(s))
	for idx, _ := range dp {
		dp[idx] = make([]bool, len(s))
	}

	start := 0
	maxLen := 0
	for i := len(s); i >= 0; i -= 1 {
		for j := i; j < len(s); j += 1 {
			if s[i] == s[j] {
				if j-i <= 1 { // 这里包含了2种情况了：a, aa
					dp[i][j] = true // 2. 初始条件
				} else {
					dp[i][j] = dp[i+1][j-1] // 3. 转移方程
				}
			}

			if dp[i][j] {
				if j-i+1 > maxLen {
					start = i
					maxLen = j - i + 1
				}
			}
		}
	}
	return s[start : start+maxLen]
}
