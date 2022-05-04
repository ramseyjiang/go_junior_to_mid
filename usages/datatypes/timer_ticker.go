package main

import (
	"fmt"
	"time"
)

func main() {
	NewTimer()
	AfterFunc()
	After()
	NewTicker()
}

// The Timer type represents a single event.
// When the Timer expires, the current time will be sent on C, unless the Timer was created by AfterFunc.
// A Timer must be created with NewTimer or AfterFunc.

func NewTimer() {
	fmt.Println("NewTimer start")
	t0 := time.NewTimer(50 * time.Millisecond)
	fmt.Println(t0.Stop()) // timer.Stop() is used to stop timer, if success, it will return true. Otherwise, it will return false.
	fmt.Println("NewTimer end")
}

// AfterFunc which is from time.AfterFunc(d Duration, f func()), it is more decoupled, and executes the relevant logic via Callback function.
func AfterFunc() {
	fmt.Println("AfterFunc start")

	i := 3
	c := make(chan bool)
	var f func()
	f = func() {
		i--
		if i >= 0 {
			time.AfterFunc(time.Second, f)
			fmt.Println("code in if", i)
		} else {
			fmt.Println("code in else", i)
			c <- true
		}
	}

	t := time.AfterFunc(2*time.Second, f)
	val := <-c

	// timer.Stop() is used to stop timer, if success, it will return true. Otherwise, it will return false.
	fmt.Println(val, ",timer.Stop() is", t.Stop())

	fmt.Println("AfterFunc end, print timer structure is", t)
}

// After which is from time.After(d Duration) <-chan Time, it is simpler.
// But problems may arise when you use two After in one select clause.
// The one that waits longer may never get the chance to execute.
// The reason is both Afters will be reset after the first After is executed,
// so the timer that waits longer will never be triggered unless you define the timer outside the loop.
func After() {
	fmt.Println("After start")
	ch := make(chan bool)

	// If comment the below anonymous func, this After() will output timeout
	go func() {
		// do something
		ch <- true // success
	}()

	select {
	case tag := <-ch:
		fmt.Println("this uses After success is", tag)
	case <-time.After(2 * time.Second): // fail safe
		fmt.Println("this After is timeout")
	}

	fmt.Println("After end")
}

// Ticker is a timer that can be executed repeatedly without a manual reset.
// A Ticker holds a channel that delivers ``ticks'' of a clock at intervals.
// Remember to initialize it separately instead of putting it into select, otherwise, it will cause memory leak.
// time.Tick() can not be collected by GC.
// So if it is not necessary, don't suggest using it.

func NewTicker() {
	fmt.Println("NewTicker start")

	ticker := time.NewTicker(2 * time.Second)
	ch := make(chan bool)

	go func() {
		// ticker.Stop()	// If ticker.Stop() is uncommented, "case tm := <-ticker.C" won't be triggered.
		time.Sleep(7 * time.Second)
		ch <- true
	}()

	for {
		select {
		case <-ch:
			fmt.Println("Completed!")
			fmt.Println("NewTicker end")
			return
		case tm := <-ticker.C: // If ticker.Stop() is uncommented, this case won't run.
			fmt.Println("The Current time is: ", tm)
		}
	}

	// fmt.Println("Ticker is turned off!")
}
