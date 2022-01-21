package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	re := bufio.NewReader(os.Stdin)
	str, _ := re.ReadString('\n')
	strSlice := strings.Split(strings.Trim(str, "\n"), " ")
	cnt := 0
	for _, val := range strSlice {
		if val == "" {
			continue
		}
		cnt += 1
	}
	fmt.Println(cnt)
}
