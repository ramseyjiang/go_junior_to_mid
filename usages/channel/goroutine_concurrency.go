package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	list := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	noCurrency(list)
	withGoroutineConcurrency(list)
}

/**
OUTPUT is below:
Current value is : a
Current value is : b
Current value is : c
Current value is : d
Current value is : e
Current value is : f
Current value is : g
Current value is : h
Program finished executing
*/
func noCurrency(list []string) {
	fmt.Println("noCurrency program begin executing")
	for i := 0; i < len(list); i++ {
		d := time.Duration(rand.Intn(10)) * time.Microsecond
		time.Sleep(d)
		fmt.Printf("Current value is : %s \n", list[i])
	}

	fmt.Println("-----noCurrency program finished executing------")
}

func withGoroutineConcurrency(list []string) {
	fmt.Println("------withGoroutineConcurrency program begin executing------")

	// WaitGroup is available in the sync package, it has functionalities that allow blocking and waiting
	// for any number of goroutines to finish executing.
	var wg sync.WaitGroup
	wg.Add(len(list))

	for i := 0; i < len(list); i++ {
		go func(val string) { // executed in a separate goroutine by using the keyword go
			// To signal the completion of the execution of a goroutine, it must call theDone method
			// on the WaitGroup object at the end of it’s execution
			defer wg.Done()
			d := time.Duration(rand.Intn(10)) * time.Microsecond
			time.Sleep(d)
			fmt.Printf("Current value is : %s \n", val)
		}(list[i])
	}

	// wg.Wait() is used to inform our current goroutine, which is the one executing the main method to wait
	// till the execution of all goroutines added using Add are finished.
	// This is done by calling the Wait method on the WaitGroup instance from the current goroutine.
	wg.Wait()
	fmt.Println("------withGoroutineConcurrency program finished executing------")
}
