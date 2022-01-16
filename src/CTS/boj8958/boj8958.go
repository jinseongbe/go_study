package main

import (
	"bufio"
	"fmt"
	"os"
)

func scoreQuiz(str string) int {
	totalScore := 0
	tempScore := 0

	for i := 0; i < len(str); i++ {

		if str[i] == 'O' {
			tempScore = tempScore + 1
			totalScore = totalScore + tempScore
		} else if str[i] == 'X' {
			tempScore = 0
		}
	}
	return totalScore
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	re := bufio.NewReader(os.Stdin)

	for i := 0; i < n; i++ {
		line, _ := re.ReadString('\n')
		result := scoreQuiz(line)
		fmt.Println(result)
	}
}
