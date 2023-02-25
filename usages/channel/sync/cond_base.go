package main

import (
	"fmt"
	"sync"
	"time"
)

// Scenario is that we have a queue of fixed length 2, and 10 items we want to push onto the queue.
// We want to enqueue items as soon as there is room, so we want to be notified as soon as there’s room in the queue.

func main() {
	// create our condition using a standard sync.Mutex as the Locker.
	c := sync.NewCond(&sync.Mutex{})

	// create a slice with a length of zero. Since we know we’ll eventually add 10 items, we instantiate it with a capacity of 10.
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		// once again enter the critical section for the condition, so we can modify data pertinent to the condition.
		c.L.Lock()

		// trigger dequeue an item by reassigning the head of the slice to the second item.
		queue = queue[1:]
		fmt.Println("Removed from queue")

		// exit the condition’s critical section since we’ve successfully dequeued an item.
		c.L.Unlock()

		// let a goroutine waiting on the condition know that something has occurred.
		// The Signal is the Cond type provides for notifying goroutines blocked on a Wait call that the condition has been triggered.
		// Signal finds the goroutine that’s been waiting the longest and notifies that.
		// Broadcast sends a signal to all goroutines that are waiting.
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		// enter the critical section for the condition by calling Lock on the condition’s Locker.
		c.L.Lock()
		for len(queue) == 2 {
			// call Wait, which will suspend the main goroutine until a signal on the condition has been sent.
			c.Wait()
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})

		// create a new goroutine that will dequeue an element after one second.
		go removeFromQueue(1 * time.Second)

		// exit the condition’s critical section since we’ve successfully enqueued an item.
		c.L.Unlock()
	}
}
