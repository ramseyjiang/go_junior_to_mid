package main

import (
	"sync"
	"testing"
)

// Benchmark is built in tool, which is used to help you monitor the source function performance.
/*
% go test -bench .
goos: darwin
goarch: arm64
pkg: github.com/ramseyjiang/go_junior_to_mid/usages/channel/chanbenchmarktest
BenchmarkWorkerProcessSetupTest1-10                  213           5632163 ns/op
BenchmarkWorkerProcessSetupTest10-10                2076            563016 ns/op
BenchmarkWorkerProcessSetupTest100-10              21482             56134 ns/op
BenchmarkWorkerProcessSetupTest1000-10            206056              5726 ns/op
BenchmarkWorkerProcessSetupTest10000-10          2252638               511.8 ns/op

% go test -bench=. -benchmem
goos: darwin
goarch: arm64
pkg: github.com/ramseyjiang/go_junior_to_mid/usages/channel/chanbenchmarktest
BenchmarkWorkerProcessSetupTest1-10                  213           5589734 ns/op               3 B/op          0 allocs/op
BenchmarkWorkerProcessSetupTest10-10                2134            558804 ns/op               3 B/op          0 allocs/op
BenchmarkWorkerProcessSetupTest100-10              21361             56298 ns/op               3 B/op          0 allocs/op
BenchmarkWorkerProcessSetupTest1000-10            202886              5635 ns/op               3 B/op          0 allocs/op
BenchmarkWorkerProcessSetupTest10000-10          2290677               511.7 ns/op             2 B/op          0 allocs/op
*/

func BenchmarkWorkerProcessSetupTest1(b *testing.B)     { benchmarkProcess(1, b) }
func BenchmarkWorkerProcessSetupTest10(b *testing.B)    { benchmarkProcess(10, b) }
func BenchmarkWorkerProcessSetupTest100(b *testing.B)   { benchmarkProcess(100, b) }
func BenchmarkWorkerProcessSetupTest1000(b *testing.B)  { benchmarkProcess(1000, b) }
func BenchmarkWorkerProcessSetupTest10000(b *testing.B) { benchmarkProcess(10000, b) }

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
