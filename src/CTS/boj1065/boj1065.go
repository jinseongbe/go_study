package main

import (
	"fmt"
)

func main() {
	var n, cnt int
	fmt.Scanf("%d", &n)

	if n >= 100 {
		cnt = 99
		for i := 100; i <= n; i++ {
			if find(i) {
				cnt = cnt + 1
			}
		}
		fmt.Println(cnt)
	} else {
		fmt.Println(n)
	}

}

func find(num int) bool {
	if num == 1000 {
		return false
	}
	first := num % 10
	second := (num / 10) % 10
	third := (num / 100) % 10

	if (third - second) == (second - first) {
		return true
	} else {
		return false
	}

}
