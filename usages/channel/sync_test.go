package channel

import (
	"sync"
	"sync/atomic"
	"testing"
)

/**
Golang synchronizes multiple goroutines via channels and sync packages to prevent multiple goroutines from competing for shared data.

% go test -bench=. -benchtime=5s
goos: darwin
goarch: arm64
pkg: junior/usages/channel
Benchmark_NoSync-10             1000000000               2.170 ns/op
Benchmark_Atomic-10             872203156                6.884 ns/op
Benchmark_Mutex-10              445907074               13.46 ns/op
Benchmark_Channel-10            227463316               26.37 ns/op
PASS
ok      junior/usages/channel   25.513s

From the results, we could identify that using channels to concurrently increase a value is much slower than the other techniques.
The atomic way is the best way.

If it is possible, don't try to not share a value between multiple goroutines,
so that it does not need to do synchronizations at all for the value.
*/

var g int32
var m sync.Mutex
var ch = make(chan struct{}, 1)

func Benchmark_NoSync(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g++
	}
}

func Benchmark_Atomic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		atomic.AddInt32(&g, 1)
	}
}

func Benchmark_Mutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m.Lock()
		g++
		m.Unlock()
	}
}

func Benchmark_Channel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch <- struct{}{}
		g++
		<-ch
	}
}
