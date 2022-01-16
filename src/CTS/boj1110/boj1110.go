package main

import (
	// "bufio"
	"fmt"
	// "os"
	// "strconv"
	// "strings"
)

func main() {
	var n, temp, a, b, cnt int
	fmt.Scanf("%d", &n)

	cnt = 0
	temp = -1

	for true {
		if n == temp {
			break
		}

		if cnt == 0 {
			temp = n
		}

		if temp < 10 { // if a+b < 10 -> b, a+b
			// fmt.Println(temp)
			a = 0
			b = temp
			// fmt.Println("case1", "a = ", a, "b = ", b)
			temp = temp*10 + (a + b)
		} else { // if a+b > 10 -> b, a+b-10
			// fmt.Println(temp)
			a = (temp / 10) % 10
			b = temp % 10
			// fmt.Println("case2", "a = ", a, "b = ", b)
			if a+b >= 10 {
				temp = b*10 + (a + b - 10)
			} else if a+b < 10 {
				temp = b*10 + (a + b)
			}
		}

		cnt = cnt + 1
	}

	fmt.Println(cnt)
}
