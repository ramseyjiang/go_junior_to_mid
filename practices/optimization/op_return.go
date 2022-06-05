package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	str0 := declareReturnTypeOnly()
	log.Println(str0)

	str1 := declareReturnNameTypeWithFmt()
	log.Println(str1)

	str2 := declareReturnNameTypeWithLog()
	log.Println(str2)
}

func declareReturnTypeOnly() string {
	execTime := time.Now()
	var allStr string
	s := [5]string{"one", "two", "3", "4", "five"}
	for _, str := range s {
		allStr += str
	}

	fmt.Printf("declareReturnTypeOnly ExecTime is %47s %6s\n", "->", time.Since(execTime))

	return allStr
}

func declareReturnNameTypeWithFmt() (allStr string) {
	execTime := time.Now()
	s := [5]string{"one", "two", "3", "4", "five"}
	for _, str := range s {
		allStr += str
	}

	fmt.Printf("declareReturnNameTypeWithFmt ExecTime is %40s %6s\n", "->", time.Since(execTime))
	return
}

func declareReturnNameTypeWithLog() (allStr string) {
	execTime := time.Now()
	s := [5]string{"one", "two", "3", "4", "five"}
	for _, str := range s {
		allStr += str
	}

	log.Printf("declareReturnNameTypeWithLog ExecTime is %20s %6s\n", "->", time.Since(execTime))
	return
}
