package main

import (
	"fmt"
	"sort"
)

func makeSelfNum(n int, numMap map[int]bool) int {
	// 맵에서 이미 있는지 판별하는 함수 사용 해서
	// 이미 있으면 return 시키고 아니면 다시 재귀함수로 계속 탐색
	switch {
	case n < 10:
		n = n + n
		if numMap[n] {
			numMap[n] = false
			makeSelfNum(n, numMap)
		} else {
			return 0
		}

	case n < 100:
		n = n + n%10 + ((n / 10) % 10)
		if numMap[n] {
			numMap[n] = false
			makeSelfNum(n, numMap)
		} else {
			return 0
		}

	case n < 1000:
		n = n + n%10 + ((n / 10) % 10) + ((n / 100) % 10)
		if numMap[n] {
			numMap[n] = false
			makeSelfNum(n, numMap)
		} else {
			return 0
		}

	case n < 10000:
		n = n + n%10 + ((n / 10) % 10) + ((n / 100) % 10) + ((n / 1000) % 10)
		if numMap[n] {
			numMap[n] = false
			makeSelfNum(n, numMap)
		} else {
			return 0
		}

	case n == 10000:
		n = 1
		if numMap[n] {
			numMap[n] = false
			makeSelfNum(n, numMap)
		} else {
			return 0
		}
	}

	return 0
}

func main() {
	numMap := make(map[int]bool)

	for i := 1; i <= 10000; i++ {
		numMap[i] = true
	}

	for i := 1; i <= 10000; i++ {
		makeSelfNum(i, numMap)
	}

	var selfNums []int
	for key, val := range numMap {
		if val {
			selfNums = append(selfNums, key)
		}
	}

	selfNums = append(selfNums, 1)

	sort.Ints(selfNums)
	for i := 0; i < len(selfNums); i++ {
		fmt.Println(selfNums[i])
	}

}

/*

공부해볼것!!!!!!!
지우지말고!!!!
package main

import (
	"bytes"
	"fmt"
	"strconv"
)

var arr [10001]bool

func d(n int) int {
	sum := n
	for n != 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func main() {
	for i := 1; i <= 10000; i++ {
		result := d(i)
		if result <= 10000 {
			arr[result] = true
		}
	}

	var buf bytes.Buffer

	arr[0] = true
	for i := range arr {
		if !arr[i] {
			buf.WriteString(strconv.Itoa(i))
			buf.WriteRune('\n')
		}
	}
	fmt.Print(buf.String())
}
*/
