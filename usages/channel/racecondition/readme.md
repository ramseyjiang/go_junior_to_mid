Two ways to run this instance.go
One way, go run instance.go,
% go run instance.go
value of j after 1000 operations is  953

The other way
% go run -race instance.go
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