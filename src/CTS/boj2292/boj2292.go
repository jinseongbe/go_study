package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	var a0 = 1
	for i := 0; i >= 0; i++ {
		an := a0 + 6*i
		if n <= an {
			fmt.Println(i + 1)
			break
		} else {
			a0 = an
		}
	}

}
