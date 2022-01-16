package main

import (
	// "bufio"
	"fmt"
	// "os"
	// "strconv"
	// "strings"
)

func makeSliceUnique1(s []string) []string {
	keys := make(map[string]struct{}) // struct는 내용이 없을시 메모리를 할당받지 않기 때문에 이렇게 사용한것!!
	res := make([]string, 0)
	for _, val := range s {
		if _, ok := keys[val]; ok {
			continue
		} else {
			keys[val] = struct{}{}
			res = append(res, val)
		}
	}
	return res
}

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
