package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scanf("%s", &str)

	strArray := [8]string{"dz=", "c=", "c-", "d-", "lj", "nj", "s=", "z="}
	var flag []int

	for i := 0; i < len(str); i++ {
		flag = append(flag, 1)
	}

	cnt := 0

	for i := 0; i < len(str)-2; i++ {
		if str[i:i+3] == strArray[0] {
			flag[i] = 0
			flag[i+1] = 0
			flag[i+2] = 0
			cnt += 1
		}
	}

	for i := 0; i < len(str)-1; i++ {
		for idx, croa := range strArray {
			if idx != 0 && str[i:i+2] == croa && flag[i] == 1 {
				flag[i] = 0
				flag[i+1] = 0
				cnt += 1
			}
		}
	}

	for _, val := range flag {
		if val == 1 {
			cnt += val
		}
	}

	fmt.Println(cnt)
}
