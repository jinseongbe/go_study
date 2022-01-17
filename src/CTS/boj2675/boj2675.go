package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	var ans []string
	fmt.Scanf("%d", &n)
	re := bufio.NewReader(os.Stdin)

	for i := 0; i < n; i++ {
		line, _ := re.ReadString('\n')
		input := strings.Split(strings.Trim(line, "\n"), " ")
		r, _ := strconv.Atoi(input[0])

		for j := 0; j < len(input[1]); j++ {
			for k := 0; k < r; k++ {
				ans = append(ans, string(input[1][j]))
			}
		}

		for _, char := range ans {
			fmt.Printf("%s", char)
		}
		ans = nil
		fmt.Println()
	}

}
