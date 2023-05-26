package niuke101b

/*
	定义4个边界，动态缩减。退出条件为边界缩减至无元素
*/
func spiralOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 {
		return nil
	}

	res := make([]int, 0)
	// 定义边界
	iLeft, jLeft := 0, 0
	iRight, jRight := len(matrix)-1, len(matrix[0])-1

	// 动态缩减
	for iLeft <= iRight && jLeft <= jRight {
		for k := jLeft; k <= jRight; k++ {
			res = append(res, matrix[iLeft][k])
		}
		iLeft++

		for k := iLeft; k <= iRight; k++ {
			res = append(res, matrix[k][jRight])
		}
		jRight--

		if iLeft <= iRight {
			for k := jRight; k >= jLeft; k-- {
				res = append(res, matrix[iRight][k])
			}
			iRight--
		}

		if jLeft <= jRight {
			for k := iRight; k >= iLeft; k-- {
				res = append(res, matrix[k][jLeft])
			}
			jLeft++
		}

	}
	return res
}
