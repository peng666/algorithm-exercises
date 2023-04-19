package main

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}
	hashTable := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if idx, ok := hashTable[target-nums[i]]; ok {
			return []int{idx, i}
		}
		hashTable[nums[i]] = i
	}
	return nil
}
