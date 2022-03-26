Mutex (mutual exclusion) is a concept in programming where only one routine (thread) can perform multiple operations at a time.
It is done by one routine acquiring the lock on value, do whatever manipulation on the value it has to do and then releasing the lock.
When the value is locked, no other routine can read or write to it.

In Go, the mutex data structure (a map) is provided by sync package. 
Before performing any operation on a value which might cause race condition, 
programs acquire a lock using mutex.Lock() method followed by code of operation. 

Once we are done with the operation, we unlock it using mutex.Unlock() method. 
When any other goroutine is trying to read or write value when the lock is present,
that goroutine will block until the operation is unlocked from the first goroutine. 

Hence, only 1 goroutine can get to read or write value, avoiding race conditions. 

Remember, any variables present in operations between the lock and unlock will not be available for other goroutines until the whole operations are unlocked.