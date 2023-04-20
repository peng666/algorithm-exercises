package main

/*

 */
func permute(nums []int) [][]int {
	result := [][]int{}
	tmp := []int{}
	visited := make([]bool, len(nums))
	var dfs func([]int)
	dfs = func(nums []int) {
		if len(nums) == len(tmp) {
			result = append(result, append([]int{}, tmp...))
			// 注意，切片实际上是一个地址（指针）result存的是各个指针
			// 而tmp是外部变量，没有重新赋值，指向的是一直不变的
			// 所以，每次都要创建一个新空间，深拷贝
		}
		for i := 0; i < len(nums); i++ {
			if visited[i] {
				continue
			}
			visited[i] = true
			tmp = append(tmp, nums[i])
			dfs(nums)
			tmp = tmp[:len(tmp)-1]
			visited[i] = false
		}
	}
	dfs(nums)
	return result
}

func main() {
	nums := []int{1, 2, 3}
	permute(nums)
}
