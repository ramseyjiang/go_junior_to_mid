package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var once = sync.Once{}

	printOnce(&once)
	fmt.Println(11)
	printOnce(&once)
	fmt.Println(22)

	time.Sleep(1) // wait for goroutines to complete
}

func printOnce(once *sync.Once) {
	// In once.Do(), you can define a func what you want to do.
	once.Do(func() {
		fmt.Println("print once")
	})
}
