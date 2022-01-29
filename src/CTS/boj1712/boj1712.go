package main

import "fmt"

func main() {
	var fixCost, varCost, price, totalCost, totalSales int
	fmt.Scanf("%d %d %d", &fixCost, &varCost, &price)

	if price <= varCost {
		fmt.Println(-1)
	} else {
		for i := 0; i >= 0; i++ {
			totalCost = fixCost + varCost*i
			totalSales = price * i
			if totalSales > totalCost {
				fmt.Println(i)
				break
			}
		}
	}
}
