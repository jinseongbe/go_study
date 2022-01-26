package main

import (
	"fmt"
)

func main() {
	var n int
	var str string

	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &str)
		if checkGroupWord(str) {
			fmt.Println(str, "is group-word")
		}
	}

}

func checkGroupWord(str string) bool {
	groupFlag := true
	alpha := "abcdefghijklmnopqrstuvwxyz"
	var flags [26]int
	for i := 0; i < 26; i++ {
		flags[i] = 0
	}

	for i := 0; i < len(str); i++ {
		for idx, char := range alpha {
			if flags[idx] != 1 && rune(str[i]) == char {
				if i != 0 && str[i-1] != str[i] && flags[idx] == 1 {
					groupFlag = false
					break
				}
				flags[idx] = 1
			}
		}

	}

	return groupFlag
}
