package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var sharedMapForMutex = map[string]int{}
var mapMutex = sync.Mutex{}

const mutexWriterRunSecond = 3

func runMapMutexReader(ctx context.Context, readChan chan int) {
	readCount := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("reader exiting...")
			readChan <- readCount
			return
		default:
			mapMutex.Lock()
			var _ = sharedMapForMutex["key"]
			mapMutex.Unlock()
			readCount += 1
		}
	}
}

// runMapMutexWriter, the writer updates the “sharedMap” every 100 milliseconds.
func runMapMutexWriter(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("writer exiting...")
			return
		default:
			mapMutex.Lock()
			sharedMapForMutex["key"] += 1
			mapMutex.Unlock()
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// main is used 15 reader goroutines and a single writer goroutine.
func main() {
	testContext, cancel := context.WithCancel(context.Background())

	readCh := make(chan int)
	sharedMapForMutex["key"] = 0

	numberOfReaders := 15
	for i := 0; i < numberOfReaders; i++ {
		go runMapMutexReader(testContext, readCh)
	}
	go runMapMutexWriter(testContext)

	// The writer updates the “sharedMap” every 100 milliseconds.
	// writeCount = writerRunSecond * time.Second * 1000 / 100
	time.Sleep(mutexWriterRunSecond * time.Second)

	cancel()

	totalReadCount := 0
	for i := 0; i < numberOfReaders; i++ {
		totalReadCount += <-readCh
	}

	time.Sleep(1 * time.Second)

	writeCount := sharedMapForMutex["key"]
	fmt.Printf("[MUTEX] Write Counter value: %v\n", writeCount)
	fmt.Printf("[MUTEX] Read Counter value: %v\n", totalReadCount)
	fmt.Println(sharedMapForMutex)
}
