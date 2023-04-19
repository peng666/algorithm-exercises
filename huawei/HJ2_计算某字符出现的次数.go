package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	chmap := make(map[byte]int)
	for input.Scan() {
		str1 := input.Text()
		input.Scan()
		str2 := input.Text()
		str1 = strings.ToLower(str1)
		str2 = strings.ToLower(str2)
		for _, c := range str1 {
			chmap[byte(c)]++
		}
		fmt.Println(chmap[str2[0]])
	}
}
