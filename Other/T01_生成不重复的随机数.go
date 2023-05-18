package main

//import (
//	"fmt"
//	"math/rand"
//)
//
///*
//	生成不重复的随机数
//	1. 去重法。记录之前生成过的随机数，然后每次生成新的都去校对。效率不高。尤其是比如从100个数生成第99个随机数时，冲突的概率很高。
//	2. 筛选法。先把所有的数装入一个数组，然后随机生成下标，然后取值并删除这个数保证下次取值不重复。实现方案就是逆序遍历，每次跟当前位交换值。
//*/
//func main() {
//	var m, n int
//	fmt.Println("请输入m和n，以空格分割，回车确认")
//	fmt.Scanf("%d %d", &m, &n)
//
//	// 初始化候选数组
//	nums := make([]int, m)
//	for i := 0; i < m; i++ {
//		nums[i] = i
//	}
//
//	// 计算随机值
//	for i := len(nums) - 1; i >= m-n; i-- {
//		idxRand := rand.Intn(i) + 1
//		nums[i], nums[idxRand] = nums[idxRand], nums[i]
//	}
//
//	// 输出结果
//	for _, num := range nums[m-n:] {
//		fmt.Printf("%v ", num)
//	}
//	fmt.Println()
//}
