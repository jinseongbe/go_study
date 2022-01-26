package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string

	var alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var index = "22233344455566677778889999"
	var nums []string

	fmt.Scanf("%s", &str)

	for _, ch := range str {
		for idx, c := range alpha {
			if ch == c {
				nums = append(nums, string(index[idx]))
			}
		}
	}

	cnt := 0
	for _, num := range nums {
		temp, _ := strconv.Atoi(num)
		cnt = cnt + (temp + 1)
	}

	fmt.Println(cnt)
}
