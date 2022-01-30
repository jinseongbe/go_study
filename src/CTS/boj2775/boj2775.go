package main

import (
	"fmt"
)

func main() {
	var T, floor, roomNum int

	fmt.Scanf("%d", &T)

	for i := 0; i < T; i++ {

		fmt.Scanf("%d\n%d", &floor, &roomNum)

		var apart [][]int
		for j := 0; j <= floor; j++ {
			var floorSlice []int
			for k := 0; k < roomNum; k++ {
				if j == 0 {
					floorSlice = append(floorSlice, k+1)
				} else {
					temp := 0
					for l := 0; l <= k; l++ {
						temp += apart[j-1][l]
					}
					floorSlice = append(floorSlice, temp)
				}
			}
			apart = append(apart, floorSlice)
		}
		fmt.Println(apart[floor][roomNum-1])

	}
}
