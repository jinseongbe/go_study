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

func main() {
	re := bufio.NewReader(os.Stdin)
	var n, x int
	fmt.Scanf("%d %d", &n, &x)
	line, _ := re.ReadString('\n')
	numStr := strings.Split(strings.Trim(line, " \n"), " ")
	numSlice := strToInt(numStr)

	for i := 0; i < n; i++ {
		if numSlice[i] < x {
			fmt.Printf("%d ", numSlice[i])

		}
	}
}
