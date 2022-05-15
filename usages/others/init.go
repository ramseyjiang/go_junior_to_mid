package main

import (
	"log"
	"time"
)

var sort = AnswerForInit()

func AnswerForInit() int8 { // 1
	log.Printf("AnswerForInit run time is %10s %s\n", "", time.Now())
	return 1
}

// init() is the go default keyword, if change it to Init(), it won't run before main()
func init() { // 2
	log.Printf("Init run time is %19s %s\n", "", time.Now())
	log.Println(sort)
	sort = 2
}

func main() { // 0
	log.Printf("Main run time is %19s %s\n", "", time.Now())

	log.Println(sort)
}
