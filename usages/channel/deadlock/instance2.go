package main

import "fmt"

func service() {
	fmt.Println("Hello from service!")
}

func main() {
	go service()

	// multichanselect will block the main goroutine,
	select {}

	fmt.Println("test")
}

/** Output is below:
Hello from service!
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [multichanselect (no cases)]:
main.main()
        /Users/daweijiang/Desktop/GoApp/go_junior/usages/channel/deadlock/instance2.go:13 +0x34
exit status 2
*/

/**
When using multichanselect, the scheduler will schedule another available goroutine which is working.
But after that, it will die and the schedule has to schedule another available goroutine,
but since main routine is blocked and no other goroutines are available, resulting in a deadlock.
*/
