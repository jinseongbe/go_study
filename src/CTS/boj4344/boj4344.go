package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func strToInt(s []string) (int, []float64, int) {
	var numSlice []float64
	var sum int
	total, _ := strconv.Atoi(s[0])
	for i := 1; i < len(s); i++ {
		num, _ := strconv.Atoi(s[i])
		sum = sum + num
		numSlice = append(numSlice, float64(num))
	}

	return total, numSlice, sum
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	re := bufio.NewReader(os.Stdin)

	for i := 0; i < n; i++ {
		line, _ := re.ReadString('\n')
		total, scores, sum := strToInt(strings.Split(strings.Trim(line, "\n"), " "))

		mean := float64(sum) / float64(total)

		var cnt int
		for j := 0; j < len(scores); j++ {
			if scores[j] > mean {
				cnt = cnt + 1
			}
		}

		fmt.Printf("%0.3f", float64(cnt)/float64(total)*100)
		fmt.Print("%\n")
	}
}
