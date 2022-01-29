// 해결못함 (unsolved-01)

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	var str string
	re := bufio.NewReader(os.Stdin)
	fmt.Fscanln(re, &n)

	cnt := 0
	for i := 0; i < n; i++ {
		fmt.Fscanln(re, &str)

		if checkGroupWord(str) {
			cnt += 1
		}
	}

	fmt.Println(cnt)
}

func checkGroupWord(str string) bool {
	alphaMap := make(map[byte]int)
	flag := true

	for i := 0; i < len(str); i++ {
		if len(str) == 1 {
			break
		}
		if i != (len(str) - 1) {
			if str[i] == str[i+1] {
				continue
			} else if str[i] != str[i+1] {
				_, exists := alphaMap[str[i]]
				if !exists {
					alphaMap[str[i]] = 0
				} else {
					flag = false
					break
				}
			}
		} else if i == (len(str) - 1) {
			if str[i] == str[i-1] {
				continue
			} else if str[i] != str[i-1] {
				_, exists := alphaMap[str[i]]
				if !exists {
					alphaMap[str[i]] = 0
				} else {
					flag = false
					break
				}
			}
		}
	}

	return flag
}
