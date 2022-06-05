package main

import (
	"log"
	"time"
)

func main() {
	s := [5]string{"one", "two", "3", "4", "five"}
	str0 := declareReturnSliceNoLength(s)
	log.Println(len(str0))

	str1 := declareReturnSliceWithLength(s)
	log.Println(len(str1))
}

func declareReturnSliceNoLength(s [5]string) (allSlice []string) {
	execTime := time.Now()
	for sum := 0; sum < 100; sum++ {
		for _, v := range s {
			allSlice = append(allSlice, v)
		}
	}

	log.Printf("declareReturnSliceNoLength ExecTime is %22s %6s\n", "->", time.Since(execTime))
	return
}

func declareReturnSliceWithLength(s [5]string) (allSlice [500]string) {
	execTime := time.Now()
	for sum := 0; sum < 100; sum++ {
		for i, v := range s {
			// allSlice = append(allSlice, v)
			allSlice[i] = v
		}
	}

	log.Printf("declareReturnSliceWithLength1 ExecTime is %20s %6s\n", "->", time.Since(execTime))
	return
}
