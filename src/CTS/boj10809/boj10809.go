package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	re := bufio.NewReader(os.Stdin)

	str, _ := re.ReadString('\n')
	var alphaArray [26]int
	for i := 0; i < 26; i++ {
		alphaArray[i] = -1
	}

	for i := 0; i < len(str); i++ {

		for j := 0; j < 26; j++ {
			if str[i] == alphabet[j] && alphaArray[j] == -1 {
				alphaArray[j] = i
			}
		}
	}

	for i := 0; i < 26; i++ {
		fmt.Printf("%d ", alphaArray[i])
	}
}
