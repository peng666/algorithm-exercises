package T4_字符串

/*
   思路：
   方法1.动态规划（没写）
   方法2.中心扩展法
       以单字母、双字母为中心点，向两边同时扩展
   方法3.马拉车算法（主要是记录下曾经计算过的最长子串半径）

*/
func Solution(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		s1 := huiwenzichuan(s, i, i)
		s2 := huiwenzichuan(s, i, i+1)
		if len(res) < len(s1) {
			res = s1
		}
		if len(res) < len(s2) {
			res = s2
		}
	}
	return res
}

func huiwenzichuan(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}
