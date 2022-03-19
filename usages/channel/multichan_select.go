package channel

import (
	"fmt"
	"time"
)

// There may be multiple channels that a function is waiting on. For this, we can use a select statement.
// speed1 sleep longer than speed2, that's why speed2 is the first arrive.

/**
The first to arrive is:
speed 2
*/

func TriggerMultiChanSelect() {
	c1 := make(chan string)
	c2 := make(chan string)
	go speed1(c1)
	go speed2(c2)
	fmt.Println("The first to arrive is:")
	select {
	case s1 := <-c1:
		fmt.Println(s1)
	case s2 := <-c2:
		fmt.Println(s2)
	}
}

func speed1(ch chan string) {
	time.Sleep(2 * time.Second) // speed1 sleep longer than speed2, that's why speed2 is the first arrive.
	ch <- "speed 1"
}

func speed2(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "speed 2"
}
