package T4_字符串

/*
思路：
	滑动窗口（快慢指针）。使用哈希表map记录是否重复，遇到重复就逐个delete不要的元素

*/

func Solution(s string) int {
	hash := make(map[byte]int)
	n := len(s)
	l := 0
	ans := 0
	for r := 0; r < n; r++ {
		if last, ok := hash[s[r]]; ok {
			for j := l; j <= last; j++ {
				delete(hash, s[j])
			}
			l = last + 1
		}
		hash[s[r]] = r
		ans = max(r-l+1, ans)
	}
	return ans
}
func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
