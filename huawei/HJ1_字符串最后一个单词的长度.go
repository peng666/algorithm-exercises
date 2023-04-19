package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	strArr := strings.Split(input.Text(), " ")

	res := len(strArr[len(strArr)-1])

	fmt.Println(res)
}
