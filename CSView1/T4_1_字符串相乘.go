//package CSView1
//
//import "strconv"
//
//func multiply(num1, num2 string) string {
//	if num1 == "0" || num2 == "0" {
//		return "0"
//	}
//
//	arr := make([]int, len(num1)+len(num2))
//	for i := len(num2) - 1; i >= 0; i-- {
//		n2 := int(num2[i] - '0')
//		for j := len(num1) - 1; j >= 0; j-- {
//			n1 := int(num1[j] - '0')
//
//			sum := n1*n2 + arr[i+j+1]
//			arr[i+j+1] = sum % 10
//			arr[i+j] += sum / 10
//		}
//	}
//	res := ""
//	for i := len(arr) - 1; i >= 0; i-- {
//		if i == 0 && arr[i] == 0 {
//			break
//		}
//		res = strconv.Itoa(arr[i]) + res
//	}
//	return res
//}
