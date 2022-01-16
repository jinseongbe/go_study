package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func strToInt(slice []string) []int {
	var intSlice []int
	for i := 0; i < len(slice); i++ {
		temp, _ := strconv.Atoi(slice[i])
		intSlice = append(intSlice, temp)
	}
	return intSlice
}

func findMinMax(numSlice []int) (int, int) {
	var min, max = 1000001, -1000001

	for i := 0; i < len(numSlice); i++ {

		if numSlice[i] < min {
			min = numSlice[i]
		}

		if numSlice[i] > max {
			max = numSlice[i]
		}
	}
	return min, max
}

func main() {

	re := bufio.NewReader(os.Stdin)

	re.ReadBytes('\n')
	line, _ := re.ReadString('\n')
	numStr := strings.Split(strings.Trim(line, " \n"), " ")
	fmt.Printf("%T\n", numStr)

	numSlice := strToInt(numStr)
	min, max := findMinMax(numSlice)

	fmt.Println(min, max)
}
