package main

import (
	"sync"
	"testing"
)

func BenchmarkWorker_ProcessSetupTest_1(b *testing.B)     { benchmarkProcess(1, b) }
func BenchmarkWorker_ProcessSetupTest_10(b *testing.B)    { benchmarkProcess(10, b) }
func BenchmarkWorker_ProcessSetupTest_100(b *testing.B)   { benchmarkProcess(100, b) }
func BenchmarkWorker_ProcessSetupTest_1000(b *testing.B)  { benchmarkProcess(1000, b) }
func BenchmarkWorker_ProcessSetupTest_10000(b *testing.B) { benchmarkProcess(10000, b) }

func benchmarkProcess(workerCount int, b *testing.B) {
	wg := &sync.WaitGroup{}
	testWorker := newWorker(workerCount, make(chan int, workerCount), wg)
	testWorker.start()

	for i := 0; i < b.N; i++ {
		testWorker.process(i)
	}

	// Stop the worker
	wg.Wait()
}
