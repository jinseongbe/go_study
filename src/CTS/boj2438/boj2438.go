package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		s += "*"
		fmt.Println(s)
	}
}
