package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Scanf("%d %d", &a, &b)

	for i := 1; i <= 100; i *= 10 {
		if ((a / i) % 10) > ((b / i) % 10) {

			fmt.Printf("%d%d%d", (a % 10), ((a / 10) % 10), ((a / 100) % 10))
			break
		} else if ((a / i) % 10) < ((b / i) % 10) {

			fmt.Printf("%d%d%d", (b % 10), ((b / 10) % 10), ((b / 100) % 10))
			break
		} else {
			continue
		}
	}

}
