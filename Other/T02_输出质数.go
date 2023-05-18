package main

import "fmt"

func main() {
	var n int
	fmt.Println("请输入数字n")
	fmt.Scanf("%d", &n)

	pri := make([]int, n+1)
	for i := 2; i <= n/i; i++ { // 如果是合数，在0~根号i之间肯定有因子，所以只需要遍历到根号i就好了
		if pri[i] == 0 {
			for j := 2 * i; j <= n; j = j + i { // 该数的倍数全部置为1合数
				pri[j] = 1 // 标记为合数
			}
		}
	}

	ret := []int{}
	for i := 2; i <= n; i++ {
		if pri[i] == 0 { // 没被标记过的就是素数
			ret = append(ret, i)
			fmt.Printf("%v ", i)
		}
	}
	fmt.Println()
}
