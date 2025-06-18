package main

func partitionLabels(s string) []int {
	hashmap := make(map[int32]int, 26)
	for i, ch := range s {
		hashmap[ch-'a'] = i
	}

	var res []int
	left, right := 0, 0
	for i, ch := range s {
		if hashmap[ch-'a'] > right {
			right = hashmap[ch-'a']
		}
		if i == right {
			res = append(res, right-left+1)
			left = right + 1
		}
	}
	return res
}
