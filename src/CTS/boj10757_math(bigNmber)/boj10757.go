package main

import (
	"fmt"
	"strconv"
)

func calculater(a string, b string, sumNums []int) []int {
	long := len(a)
	short := len(b)
	tempUp := 0
	for i := 0; i < short; i++ {
		tempA, _ := strconv.Atoi(string(a[long-i-1]))
		tempB, _ := strconv.Atoi(string(b[short-i-1]))
		tempSum := tempA + tempB + tempUp
		sumNums = append(sumNums, tempSum%10)
		if tempSum < 10 {
			tempUp = 0
		} else {
			tempUp = 1
		}
	}

	if len(a) != len(b) {
		for j := short; j < long; j++ {
			tempA, _ := strconv.Atoi(string(a[long-j-1]))
			tempSum := tempA + tempUp
			sumNums = append(sumNums, tempSum%10)
			if tempSum < 10 {
				tempUp = 0
			} else {
				tempUp = 1
			}
			if j == long-1 && tempUp == 1 {
				sumNums = append(sumNums, 1)
			}
		}
	} else {
		if tempUp == 1 {
			sumNums = append(sumNums, 1)
		}
	}

	return sumNums
}

func main() {
	var a, b string
	var sumNums []int

	fmt.Scanf("%s %s", &a, &b)

	if len(a) > len(b) {
		sumNums = calculater(a, b, sumNums)
	} else if len(a) < len(b) {
		sumNums = calculater(b, a, sumNums)
	} else {
		sumNums = calculater(a, b, sumNums)
	}

	for i := 0; i < len(sumNums); i++ {
		fmt.Print(sumNums[len(sumNums)-i-1])
	}
}

// 라이브러리 사용
/*
package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

var in = bufio.NewReader(os.Stdin)

func main() {
	a, b := new(big.Int), new(big.Int)
	fmt.Fscan(in, a, b)
	fmt.Println(a.Add(a, b))
}
*/
