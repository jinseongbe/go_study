package main

import (
	// "bufio"
	"fmt"
	// "os"
	"strconv"
	// "strings"
)

func main() {
	var a, b, c, num int
	numIndex := "0123456789"

	fmt.Scanf("%d\n%d\n%d", &a, &b, &c)

	num = a * b * c
	numStr := strconv.Itoa(num)

	// fmt.Println(num)

	for j := 0; j < 10; j++ {
		cnt := 0
		for i := 0; i < len(numStr); i++ {
			if numStr[i] == numIndex[j] {
				cnt = cnt + 1
			}
		}
		fmt.Println(cnt)
	}
}
