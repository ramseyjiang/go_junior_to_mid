**What is the Channel?**

Channels are a mechanism that allow goroutines (which are similar to threads) to pass data from one to another. 
That makes it easier to write thread-safe concurrent programs and prevent race conditions.

It provides go keyword to create goroutines. When go keyword is placed before a function call, it becomes goroutines.
Goroutines behave like threads but technically; it is an abstraction over threads.

When we run a Go program, Go runtime will create few threads on a core on which all the goroutines are multiplexed (spawned). 
At any point in time, one thread will be executing one goroutine and if that goroutine is blocked, 
then it will be swapped out for another goroutine that will execute on that thread instead. 
This is like thread scheduling but handled by Go runtime and this is much faster.

Other languages, the traditional way to pass data from one thread to another is through shared memory that is safe-guarded by a lock.

One thread writes data to memory and another thread reads or updates the data from that same spot in memory. 

But threads are unpredictable and could end up conflicting with one-another, which is why a lock is needed.
However, using a lock can be very tricky to do correctly. 
It’s the source of a litany of hard-to-detect bugs. Channels allow you to avoid those bugs by abstracting away that complexity into a simple-to-use interface.

Closing a channel prevents a goroutine from sending data into it. But it does not prevent consuming data from the channel. 
This makes sense as we would not want to risk crucial data being left orphaned.

In Go, receiving data from an empty channel operates similar to sending data to a full channel, but with a slight twist.

Whenever a goroutine attempts to pull data from an empty channel, it will block until the channel has data.
The goroutine is added to a queue in channels struct for receivers.
But when another goroutine attempts to send data to a channel with waiting receivers, that data does not get added to the buffer.
Instead, the runtime is passes the data directly to the first waiting goroutine, bypassing the buffer entirely. 


**Goroutines**

The minimal size of a goroutine object is 2 KB, it can extend to 1 GB.
It is advised in most of the cases, to run all your goroutines on one core 
but if you need to divide goroutines among available CPU cores of your system, 
you use GOMAXPROCS environment variable or call to runtime using function runtime.GOMAXPROCS(n) where n is the number of cores to use. 

But you may sometime feel that setting GOMAXPROCS > 1 is making your program slower.

Go has an M:N scheduler that can also utilize multiple processors. 
At any time, M goroutines need to be scheduled on N OS threads that run on at most on GOMAXPROCS numbers of processors. 
At any time, at most only one thread is allowed to run per core.

If your program doesn’t start any additional goroutines, 
it will naturally run in only one thread no matter how many cores you allow it to use.!
Goroutines require less memory than threads, less time to start up and stop.
main() is a special goroutine.


**how block Goroutines and channels** 
time.Sleep can be used to block a goroutine. Channel operations are also blocking in nature. 
When some data is written to the channel, goroutine is blocked until some other goroutine reads it from that channel. 
At the same time, channel operations tell the scheduler to schedule another goroutine, 
that’s why a program does n’ot block forever on the same goroutine. 
These features of a channel are very useful in goroutines communication as it prevents us from writing manual locks and hacks to make them work with each other.


**Some differences between thread and goroutine.**

1. OS threads are managed by kernal and has hardware dependencies. 
Goroutines are managed by go runtime and has no hardware dependencies.

2. OS threads generally have fixed stack size of 1-2MB. 
Goroutines typically have 8KB (2KB since Go 1.4) of stack size in newer versions of go.

3. Stack size is determined during compile time and can not grow. 
Stack size of go is managed in run-time and can grow up to 1GB which is possible by allocating and freeing heap storage

4. There is no easy communication medium between threads. There is huge latency between inter-thread communication. 
Goroutine use channels to communicate with other goroutines with low latency.

5. Threads have identity. There is TID which identifies each thread in a process. 
Goroutine do not have any identity. Go implemented this because go does not have TLS(Thread Local Storage).

6. Threads have significant setup and teardown cost as a thread has to request lots of resources from OS and return once it's done.	Goroutines are created and destroyed by the go's runtime. These operations are very cheap compared to thread as go runtime already maintain pool of threads for goroutines. In this case OS is not aware of goroutines.

7. Threads are preemptively scheduled. Switching cost between threads is high as scheduler needs to save/restore more than 50 registers and states. This can be quite significant when there is rapid switching between threads.	
Goroutines are coopertively scheduled (read more). When a goroutine switch occurs, only 3 registers need to be saved or restored.