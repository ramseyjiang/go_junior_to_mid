package main

import (
	"fmt"
	"time"
)

/**
select statement is like switch but instead of boolean operations, choose channel operations read or write or mixed of read and write.
The select statement is blocking except when it has a default case.

If all case statements are blocking then select statement will wait until one of them unblocks and that case will be executed.
If some or all of the channel operations are non-blocking, then one of them will be chosen randomly and executed immediately.
*/

func main() {
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
