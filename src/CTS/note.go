package main

import (
	"fmt"
)

func main() {
	var slice1 [][]int

	for j := 0; j < 5; j++ {
		var slice2 []int
		for i := 0; i < 10; i++ {
			slice2 = append(slice2, (2*i)+(j-8))
		}
		slice1 = append(slice1, slice2)
	}

	for _, val := range slice1 {
		fmt.Println(val)
	}
	var in, out int
	fmt.Scanf("%d %d", &in, &out)
	fmt.Println("in : ", in, "out : ", out)

	fmt.Printf("slice1[%d][%d]: %d", in, out, slice1[in][out])

}
