package main

import (
	"fmt"
	"time"
)

var sort = AnswerForInit()

func AnswerForInit() int { // 1
	fmt.Println("AnswerForInit run time is", time.Now())
	return 42
}

func init() { // 2
	fmt.Println("Init run time is", time.Now())
	sort = 0
}

func main() { // 3
	fmt.Println("Main run time is", time.Now())
	if sort == 0 {
		fmt.Println("It's all a lie.")
	}
}
