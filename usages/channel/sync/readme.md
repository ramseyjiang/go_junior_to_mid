The sync package contains the concurrency primitives that are most useful for low level memory access synchronization.

Mutex (mutual exclusion) is a concept in programming where only one routine (thread) can perform multiple operations at a time.
It is done by one routine acquiring the lock on value, do whatever manipulation on the value it has to do and then releasing the lock.
When the value is locked, no other routine can read or write to it.

In Go, the mutex data structure (a map) is provided by sync package. 
Before performing any operation on a value which might cause race condition, 
programs acquire a lock using mutex.Lock() method followed by code of operation. 

Once we are done with the operation, we unlock it using mutex.Unlock() method. 
When any other goroutine is trying to read or write value when the lock is present,
that goroutine will block until the operation is unlocked from the first goroutine. 

sync.RWMutex provides two locking modes:
Lock/Unlock: Locks/unlocks the data structure in writing mode. Neither writers nor readers will be able to access it.
RLock/RUnlock: Locks/unlocks the data structure in reading mode. Readers will be able to access it, but not writers. 
Using this, you can achieve better performance in scenarios with a lot of readers and few writers.

Hence, only 1 goroutine can get to read or write value, avoiding race conditions. 

Remember, any variables present in operations between the lock and unlock will not be available for other goroutines until the whole operations are unlocked.

"Do not communicate by sharing memory; instead, share memory by communicating."
This means that instead of struggling with complex mutex situations in shared memory, use channels to communicate goroutines.


1. sync.atomic: Useful when making operations with integers.

2. sync.Mutex and sync.RWMutex: For synchronizing the access to more complex data structures. 
It’s the classical approach and allows for custom locking.

3. channels: When mutexes are not an option, or it is complicated to operate with them.
