package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	var a, b, c, d, re int

	fmt.Println("자유롭게 숫자 네개 입력")
	re, _ = fmt.Scan(&a, &b, &c, &d)
	fmt.Printf("네수의 합은: %d\n", a+b+c+d)
	fmt.Println("re:", re)

	a, b = 0, 0
	fmt.Println("두수를 .으로 구분하여 입력(공백x)")
	re, _ = fmt.Scanf("%d.%d", &a, &b)
	fmt.Printf("두수의 합은 : %d\n", a+b)
	fmt.Println("re: ", re)

	var aa, bb = 0, 0

	fmt.Println("두수를 공백을 통해 입력")
	re, _ = fmt.Scanln(&aa, &bb)
	fmt.Printf("두수의 합은 %d\n", aa+bb)
	fmt.Println("re : ", re)

	// fmt.Fscan : bufio를 통해 빠른 입출력!
	Big := make([]int, 100000, 100000)
	fmt.Println("그냥 scan")
	start := time.Now()
	for i := 0; i < 100000; i++ {
		fmt.Println(i)
		fmt.Scan(Big[i])
	}

	fmt.Println("fmt.Scan: %v\n", time.Since(start))

	start = time.Now()
	r := bufio.NewReader(os.Stdin)
	for i := 0; i < 100000; i++ {
		fmt.Fscan(r, Big[i])
	}

	fmt.Printf("fmt.Fscan: %v\n", time.Since(start))
	fmt.Println("end!!")

}
