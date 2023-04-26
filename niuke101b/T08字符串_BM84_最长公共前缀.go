package niuke101b

/*
	分治法。递归。注意结束条件
*/
func longestCommonPrefix(strs []string) string {
	if strs == nil {
		return ""
	}
	m := len(strs)
	if m == 1 {
		return strs[0]
	}

	preLeft := longestCommonPrefix(strs[:m/2])
	preRight := longestCommonPrefix(strs[m/2:])
	idx := -1
	for i, j := 0, 0; i < len(preLeft) && j < len(preRight); {
		if preLeft[i] == preRight[j] {
			i++
			j++
			idx++
		} else {
			break
		}
	}
	return preLeft[:idx+1]
}
