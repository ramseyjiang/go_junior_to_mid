package channel

import (
	"sync"
	"testing"
)

/**
sync.Pool is a component under the sync package to create a self-managed temporary retrieval object pool.

Frequent allocation and recycling of memory will cause a heavy burden to GC.
sync.Pool can cache objects that are not used temporarily and use them directly (without reallocation) when they are needed next time.
This can potentially reduce the GC workload and improve performance.
*/

type Person struct {
	Age int
}

var personPool = sync.Pool{
	New: func() interface{} { return new(Person) },
}

func BenchmarkWithoutPool(b *testing.B) {
	var p *Person
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			p = new(Person)
			p.Age = 23
		}
	}
}

func BenchmarkWithPool(b *testing.B) {
	var p *Person
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			p = personPool.Get().(*Person)
			p.Age = 23
			personPool.Put(p)
		}
	}
}

func BenchmarkPool(b *testing.B) {
	var p sync.Pool
	b.RunParallel(func(pb *testing.PB) {
		a := 1
		for pb.Next() {
			p.Put(&a)
			p.Get()
		}
	})
}

func BenchmarkAllocation(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// i := 0
			// i = i
		}
	})
}

/**
How does sync.Pool work?
sync.Pool has two containers for objects: local pool (active) and victim cache (archived).
According to the sync/pool.go , package init function registers to the runtime as a method to clean the pools.
This method will be triggered by the GC.

func init() {
   runtime_registerPoolCleanup(poolCleanup)
}

When the GC is triggered, objects inside the victim cache will be collected and then objects inside the local pool will
be moved to the victim cache.

func poolCleanup() {
   // Drop victim caches from all pools.
   for _, p := range oldPools {
      p.victim = nil
      p.victimSize = 0
   }

   // Move primary cache to victim cache.
   for _, p := range allPools {
      p.victim = p.local
      p.victimSize = p.localSize
      p.local = nil
      p.localSize = 0
   }

   oldPools, allPools = allPools, nil
}
New objects are put in the local pool.
Calling Put method will put the object into the local pool as well.
Calling Get method will take an object from the victim cache in the first place and if the victim cache was empty the
object will be taken from the local pool.
*/
