package main

import (
	"fmt"
	"sync"
)

// Pool is a concurrent-safe implementation of the object pool pattern.

// Pool’s primary interface is its Get method.
// When called, Get will first check whether there are any available instances within the pool to return to the caller,
// and if not, call ts New member variable to create a new one.
// When finished, callers call Put to place the instance they were working with back in the pool for use by other processes.

func main() {
	basic()
	withPool()
}

func basic() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}

	// call Get on the pool, the two myPool.Get() will print "Creating new instance."
	// Because these calls will invoke the New function defined on the pool since instances haven’t yet been instantiated.
	fmt.Println("First Get")
	myPool.Get()

	fmt.Println("Second Get")
	instance1 := myPool.Get()

	// put an instance previously retrieved back in the pool, this increases the available number of instances to one.
	fmt.Println("Put")
	myPool.Put(instance1)

	// Because the previous myPool.Put(instance), the third myPool.Get() can reuse the instance in the pool. The New function will not be invoked.
	myPool.Get()
	// When the fourth myPool.Get() executes, the pool is empty, so it has to create a new instance again. That's why the New is invoked again.
	instance2 := myPool.Get()

	// put an instance previously retrieved back in the pool, this increases the available number of instances to one.
	fmt.Println("Put")
	myPool.Put(instance2)

	// It won't output "Creating new instance.", because it will use the instance in the pool
	myPool.Get()
}

// withPool uses sync.Pool, it allocates a number of KB at the end.
// Maybe 10KB, maybe 12KB, the numCalcsCreated results are nondeterministic.
// If it does not use the sync.Pool, in the worst case, it will attempt to allocate 1 GB memory.
func withPool() {
	var numCalcsCreated int
	calcPool := &sync.Pool{
		New: func() interface{} {
			numCalcsCreated += 1
			// creates a new slice of bytes with a length of 1024, which means that it has enough space to hold 1024 bytes of data.
			// Hence, the mem variable maximum size is 1KB.
			mem := make([]byte, 1024)
			// store the address of the slice of bytes
			return &mem
		},
	}

	// Seed the pool with 4KB
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()
			// assert the type is a pointer to a slice of bytes, it is the force type convert.
			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)
		}()
	}

	wg.Wait()
	fmt.Printf("%d calculators were created.", numCalcsCreated)
}
