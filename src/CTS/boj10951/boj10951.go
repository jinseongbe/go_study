package main

import (
	// "bufio"
	"fmt"
	"io"
	// "os"
	// "strconv"
	// "strings"
)

func main() {
	n := 0
	var a, b int
	for n < 1 {
		_, err := fmt.Scanf("%d %d", &a, &b)
		if err == io.EOF {
			break
		}
		fmt.Println(a + b)
	}
}
