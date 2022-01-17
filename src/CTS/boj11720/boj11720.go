package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	re := bufio.NewReader(os.Stdin)

	n, _ := re.ReadString('\n')
	num, _ := strconv.Atoi(strings.Trim(n, "\n"))

	line, _ := re.ReadString('\n')
	// fmt.Println("num =", num)

	sum := 0
	for i := 0; i < num; i++ {
		// fmt.Println(string(line[i]))
		each, _ := strconv.Atoi(string(line[i]))

		sum += each
	}

	fmt.Println(sum)
}
