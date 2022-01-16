package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func strToInt(slice []string) ([]int, int) {
	var intSlice []int
	max := -1
	for i := 0; i < len(slice); i++ {
		temp, _ := strconv.Atoi(slice[i])
		intSlice = append(intSlice, temp)

		if temp > max {
			max = temp
		}
	}
	return intSlice, max
}

func main() {
	re := bufio.NewReader(os.Stdin)
	re.ReadString('\n')
	line, _ := re.ReadString('\n')

	scores, max := strToInt(strings.Split(strings.Trim(line, "\n"), " "))
	fmt.Println(scores, max)

	sum := 0.0
	for i := 0; i < len(scores); i++ {
		sum = sum + ((float64(scores[i]) / float64(max)) * 100)
	}

	mean := sum / float64(len(scores))

	fmt.Println(mean)
}
