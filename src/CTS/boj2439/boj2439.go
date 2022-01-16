package main

import "fmt"

func main() {
	var n int
	var s1, s2 string

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			s1 += " "
		}
		s2 += "*"
		fmt.Printf("%s%s\n", s1, s2)
		s1 = ""
	}
}

// 4ms 답
//package main
//
//import (
//    "bufio"
//    "os"
//    "strconv"
//)
//
//func main() {
//    br := bufio.NewReaderSize(os.Stdin, 5)
//
//    line, _, _ := br.ReadLine()
//    n, _ := strconv.Atoi(string(line))
//
//    bw := bufio.NewWriterSize(os.Stdout, 5*n*n/2)
//    defer bw.Flush()
//
//    arrByte := make([]byte, n)
//
//    for i := 0; i < n; i++ {
//        arrByte[i] = ' '
//    }
//
//    for i := n-1; i >= 0; i-- {
//        arrByte[i] = '*'
//        bw.WriteString(string(arrByte))
//        bw.WriteByte('\n')
//    }
//}
