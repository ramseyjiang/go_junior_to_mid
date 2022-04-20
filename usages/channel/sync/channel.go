package main

import (
	"log"
	"sync"
)

// All the operations have a no-buffer res channel to receive the operation’s response and synchronize goroutines.
// This strategy is not based on sharing memory, it communicates relies on sending operations through channels.
type op struct {
	res chan int
}

type incrementOp struct {
	op
}

type getValueOp struct {
	op
}

func newIncrementOp() incrementOp {
	return incrementOp{
		op: op{
			res: make(chan int),
		},
	}
}

func newGetValueOp() getValueOp {
	return getValueOp{
		op: op{
			res: make(chan int),
		},
	}
}

// newIncrement is used FanInOut pattern.
func newIncrement(ops chan<- incrementOp, wg *sync.WaitGroup) {
	defer wg.Done()

	newAddOp := newIncrementOp()
	ops <- newAddOp // newAddOp chan is paused, because it is full.
	<-newAddOp.res  // let newAddOp chan empty, otherwise it will block newAddOp channel.
}

func main() {
	incrementOps := make(chan incrementOp)
	getValueOps := make(chan getValueOp)

	go func() {
		newCounter := 0
		for {
			select {
			case newOp := <-incrementOps:
				newCounter++
				newOp.res <- newCounter
			case newOp := <-getValueOps:
				newOp.res <- newCounter
			}
		}
	}()

	wg := sync.WaitGroup{}

	for i := 0; i < 50; i++ {
		wg.Add(2)

		go newIncrement(incrementOps, &wg)
		go newIncrement(incrementOps, &wg)
	}

	wg.Wait()

	getValueOperation := newGetValueOp()
	getValueOps <- getValueOperation // getValueOperation chan is paused, because it is full.

	log.Printf("Counter: %d", <-getValueOperation.res) // let getValueOperation chan empty.

	// If uncomment below line, it will have an error. Because getValueOperation channel is empty after line 78 is executed.
	// log.Printf("Counter: %d", <-getValueOperation.res)
}
