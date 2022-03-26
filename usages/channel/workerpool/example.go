package main

import (
	"fmt"
	"sync"
	"time"
)

const capacity = 10

// This example explains how multiple goroutines can feed on the same channel and get the job done elegantly.
// sqrWorker is a worker to make squares. Parameters are two channels and an int.
// The job of this goroutine is to send squares of the number received from tasks channel to result channel.
func sqrWorker(wg *sync.WaitGroup, tasks <-chan int, results chan<- int, instance int) {
	for num := range tasks {
		// simulating blocking task operation in worker goroutine,
		// that will call the scheduler to schedule another available goroutine until it becomes available.
		// If remove, time.Sleep() call, then only one goroutine will perform the job as no other goroutines are scheduled
		// until for range loop is done and goroutine dies.
		time.Sleep(time.Millisecond)

		fmt.Printf("[worker %v] Sending result by worker %v\n", instance, instance)

		// When worker goroutine becomes available, it writes to the results channel.
		// Writing to buffered channel is non-blocking until the buffer is full, so here writing to result channel is non-blocking.
		// Current worker goroutine was unavailable, multiple other worker goroutines were executed consuming values in tasks buffer.
		results <- num * num
	}

	// done with worker
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	// Created tasks and result channel with buffer capacity 10.
	// Any send operation will be non-blocking until the buffer is full.
	tasks := make(chan int, capacity)
	results := make(chan int, capacity)

	// Launching 3 worker goroutines.
	// Spawn multiple instances of sqrWorker as goroutines to get information on which worker is executing a task.
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go sqrWorker(&wg, tasks, results, i)
	}

	// passing 5 tasks to the tasks channel which was non-blocking
	for i := 0; i < 5; i++ {
		tasks <- i * 2 // non-blocking as buffer capacity is 10
	}

	fmt.Println("[main] Wrote 5 tasks")

	// Since done with tasks channel, close it. It will save a lot of time in the future if some bugs get in.
	close(tasks)

	// wait until all workers done their job
	// Using waitGroup, we can prevent lots of unnecessary context switching scheduling.
	// But there is a sacrifice, as you have to wait until all jobs are done.
	wg.Wait()

	// After all worker goroutines consumed tasks, for range loop finishes when tasks channel buffer is empty.
	for i := 0; i < 5; i++ {
		// pull data from results channel.
		// Since read operation on an empty buffer is blocking, a goroutine will be scheduled from the worker pool.
		// Until that goroutine returns some result, the main goroutine will be blocked.
		result := <-results
		fmt.Println("[main] Result", i, ":", result)
	}

	// After all worker goroutines died or sleeping, main goroutine will regain control and continue its execution.
	fmt.Println("main regain control")
}
