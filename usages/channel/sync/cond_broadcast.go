package main

import (
	"fmt"
	"sync"
)

func main() {
	// define a type Button that contains a condition, Clicked
	type Button struct {
		Clicked *sync.Cond
	}
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	// define a convenience function that will allow registering functions to handle signals from a condition.
	// Each handler is run on its own goroutine, and subscribe will not exit until that goroutine is confirmed to be running.
	subscribe := func(c *sync.Cond, fn func()) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}

	// Set a handler for when the mouse button is raised and created a WaitGroup.
	// It in turn calls Broad cast on the Clicked Cond to let all handlers know that the mouse button has been clicked
	var clickRegistered sync.WaitGroup

	// This is done only to ensure our program doesn’t exit before our writes to stdout occur.
	clickRegistered.Add(3)

	// register a handler that simulates maximizing the button’s window when the button is clicked.
	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing window.")
		clickRegistered.Done()
	})

	// register a handler that simulates displaying a dialog box when the mouse is clicked.
	subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialogue box!")
		clickRegistered.Done()
	})

	// register a handler that simulates other button clicked, here you can add any click event you want.
	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked.")
		clickRegistered.Done()
	})

	// After button.Clicked.Broadcast() is called, and then, all three handlers would be invoked.
	// This is something channels can’t do easily and thus is one of the main reasons to utilize the Cond type.
	button.Clicked.Broadcast()

	clickRegistered.Wait()
}
