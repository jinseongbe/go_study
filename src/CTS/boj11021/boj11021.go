package main

import (
	"fmt"
)

func main() {

	var n int
	var a, b int

	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Scan(&a, &b)
		fmt.Printf("Case #%d: %d\n", i, a+b)
	}
}
