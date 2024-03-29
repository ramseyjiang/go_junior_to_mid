What is the race condition?

When two processes try to access the same resource at the same time with one or more writing operations, both of them process writing or one process writing while the other reads, involved at that precise moment.

Go has a very handy tool to help diagnose race conditions, that you can run in your tests or your application directly.

"-race", in Golang, it is called Race detector.

Three ways to run this instance.go
One way, go run instance.go,
% go run instance.go
value of j after 1000 operations is  953

The second way
% go run -race instance.go

The third way
% go run -race .

**If code occasionally accesses shared variables, it might not be able to detect the race condition. 
To detect it, the code should run in heavy load, and race conditions must be occurring.**

==================
WARNING: DATA RACE
Read at 0x000100ce55e0 by goroutine 8:
main.worker()
/Users/daweijiang/Desktop/GoApp/go_junior/usages/channel/racecondition/instance.go:12 +0x34

Previous write at 0x000100ce55e0 by goroutine 7:
main.worker()
/Users/daweijiang/Desktop/GoApp/go_junior/usages/channel/racecondition/instance.go:12 +0x4c

Goroutine 8 (running) created at:
main.main()
/Users/daweijiang/Desktop/GoApp/go_junior/usages/channel/racecondition/instance.go:24 +0x94

Goroutine 7 (finished) created at:
main.main()
/Users/daweijiang/Desktop/GoApp/go_junior/usages/channel/racecondition/instance.go:24 +0x94
==================

value of j after 1000 operations is  1000
Found 1 data race(s)
exit status 66
