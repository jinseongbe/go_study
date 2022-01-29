package main

import "fmt"

func main() {
	var a, b, v float64
	fmt.Scanf("%f %f %f", &a, &b, &v)
	result := float64(v-b) / (a - b)

	if result == float64(int(result)) {
		fmt.Println(int(result))
	} else {
		fmt.Println(int(result) + 1)
	}
}
