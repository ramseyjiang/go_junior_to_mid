package main

// The panic function expects a single argument that satisfies the empty interface.
// That argument is converted to a string (if necessary) and printed as part of the panic’s log message.

// When a program panics, a stack trace, or listing of the call stack, is included in the panic output.
func main() {
	one()
}

func one() {
	two()
}

func two() {
	three()
}

/**
Code results are as below. The stack trace includes the list of function calls that have been made.

panic: oh, no, the program is stopped. This call stack's too deep for me

goroutine 1 [running]:
main.three(...)
        /Users/daweijiang/go/src/apps/learn/base/usages/panic_ex.go:18
main.two(...)
        /Users/daweijiang/go/src/apps/learn/base/usages/panic_ex.go:14
main.one(...)
        /Users/daweijiang/go/src/apps/learn/base/usages/panic_ex.go:10
main.main()
        /Users/daweijiang/go/src/apps/learn/base/usages/panic_ex.go:6 +0x29
exit status 2
*/
func three() {
	panic("oh, no, the program is stopped. This call stack's too deep for me")
}
