package main

import (
	"fmt"
	"strconv"
)

func main() {
	var caseNum, h, w, n int

	fmt.Scanf("%d", &caseNum)

	for i := 0; i < caseNum; i++ {
		fmt.Scanf("%d %d %d", &h, &w, &n)

		roomNum := 1
		for j := 0; j >= 0; j++ {
			if n <= h {
				break
			}
			n = n - h
			roomNum += 1
		}

		var roomNumChar string
		if roomNum < 10 {
			roomNumChar = "0" + strconv.Itoa(roomNum)
		} else {
			roomNumChar = strconv.Itoa(roomNum)
		}
		ans := strconv.Itoa(n) + roomNumChar
		fmt.Println(ans)
	}

}

// 참고용
/*
package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var T int
	if sc.Scan() {
		T, _ = strconv.Atoi(sc.Text())
	}
	for i := 0; i < T; i++ {
		if sc.Scan() {
			tmp := strings.Fields(sc.Text())

			H, _ := strconv.Atoi(tmp[0])
			//W, _ := strconv.Atoi(tmp[1])
			N, _ := strconv.Atoi(tmp[2])

			w := (N-1)/H + 1
			h := (N-1)%H + 1

			wr.WriteString(strconv.Itoa(h))
			if w < 10 {
				wr.WriteString("0")
			}
			wr.WriteString(strconv.Itoa(w))
			wr.WriteString("\n")
		}
	}
}
*/
