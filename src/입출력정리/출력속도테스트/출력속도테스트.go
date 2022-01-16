package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	var num int
	//	r := bufio.NewReader(os.Stdin)
	//	w := bufio.NewWriter(os.Stdout)

	fmt.Scan(&num)

	t1 := time.Now()
	for i := num; i >= 1; i-- {
		fmt.Println(i)
	}
	t2 := time.Since(t1)
	fmt.Printf("took %v\n", t2)

	t1 = time.Now()
	var ans string
	ans = ""

	fmt.Scan(&num)

	for i := num; i >= 1; i-- {
		ans += strconv.Itoa(i) + "\n"
	}
	fmt.Printf(ans)
	t3 := time.Since(t1)
	fmt.Printf("took %v\n", t3)

	fmt.Scan(&num)
	t1 = time.Now()
	s := make([]string, 0)

	for i := num; i >= 1; i-- {
		s = append(s, strconv.Itoa(i))
	}
	fmt.Println(strings.Join(s, "\n"))
	t5 := time.Since(t1)
	fmt.Printf("took %v\n", t5)

	fmt.Println("result")
	fmt.Println("first:", t2)
	fmt.Println("second", t3)
	fmt.Println("third:", t5)
}
