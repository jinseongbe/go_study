package main

import (
	"bufio"
	"fmt"
	"os"
)

func changeToCapital(str string) string {
	small := "abcdefghijklmnopqrstuvwxyz"
	big := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var newStr []byte

	for _, char := range str {
		for i, smallAlpha := range small {
			if char == smallAlpha {
				newStr = append(newStr, big[i])
				break
			} else if char == rune(big[i]) {
				newStr = append(newStr, big[i])
				break
			}
		}
	}
	return string(newStr)
}

func findMaxValue(s []int) int {
	max := -1
	for _, val := range s {
		if val > max {
			max = val
		}
	}

	return max
}

func main() {
	re := bufio.NewReader(os.Stdin)
	str, _ := re.ReadString('\n')

	newStr := changeToCapital(str)
	indexMap := make(map[rune]int)

	for _, runn := range newStr {
		_, exist := indexMap[runn]
		if exist {
			indexMap[runn] += 1
		} else {
			indexMap[runn] = 1
		}
	}

	var vals []int
	for _, val := range indexMap {
		vals = append(vals, val)
	}
	max := findMaxValue(vals)

	var result []rune
	for key, val := range indexMap {
		if val == max {
			result = append(result, key)
		}
	}

	if len(result) == 1 {
		fmt.Println(string(result[0]))
	} else {
		fmt.Println("?")
	}

}
