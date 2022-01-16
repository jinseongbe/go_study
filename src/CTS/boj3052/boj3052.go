package main

import (
	// "bufio"
	"fmt"
	// "os"
	// "strconv"
	// "strings"
)

func makeSliceUnique(s []int) []int {
	keys := make(map[int]struct{})
	res := make([]int, 0)

	for _, val := range s {
		if _, exist := keys[val]; exist {
			continue
		} else {
			keys[val] = struct{}{}
			res = append(res, val)
		}
	}
	return res
}

func main() {
	var nums []int
	var temp int

	for i := 0; i < 10; i++ {
		fmt.Scanf("%d", &temp)
		nums = append(nums, temp%42)
	}
	fmt.Println(nums)

	result := makeSliceUnique(nums)

	fmt.Println(result)
	fmt.Println(len(result))
}

// 참고
/*
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {
	var a, count int
	n := make([]int, 10)

	for i := 0; i < 10; i++ {
		fmt.Fscan(r, &a)
		n[i] = a % 42
	}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i != j || n[j] == -1 {
				if n[i] == n[j] {
					n[j] = -1
				}
			}
		}
	}
	for _, i := range n {
		if i != -1 {
			count++
		}
	}

	w.WriteString(strconv.Itoa(count))
	w.Flush()
}

*/
