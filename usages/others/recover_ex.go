package main

import "fmt"

/*
Go offers a built-in recover function that can stop a program from panicking.
We’ll need to use it to exit the program gracefully.

If there is no panic, calls to recover return nil.
But when there is a panic, recover returns whatever value was passed to panic.
It can be used to gather information about the panic, to aid in recovering or to report errors to the user.

The same with panic, the type for recover's return value is also interface{}.
*/

/**
If all codes below does not have the recover, it only has panic(), the result is under below.
Result:

panic: Program panic is triggered.

goroutine 1 [running]:
main.freakOut()
        /Users/daweijiang/go/src/apps/learn/base/usages/recover_ex.go:18 +0x49
main.main()
        /Users/daweijiang/go/src/apps/learn/base/usages/recover_ex.go:30 +0x19
exit status 2
*/

func freakOut(panicValue string) {
	defer calmDown()  // Defer a call to the function recover().
	panic(panicValue) // Panic() is called after defer recover().
	// recover() //If call recover() in the same function with panic(), the recover() won't work.
}

// Call recover() must be not in the same function with panic().
func calmDown() {
	// Call recover() and print the panic value.
	// recover()	//If only has recover() at here, don't use fmt.Println() it won't print the panicValue.
	fmt.Println(recover()) // The value printed is the panicValue from panic().
}

func main() {
	panicValue := "Program panic is triggered."
	exitValue := "Existing normally, the panic is not output, because recover is used."

	// fmt.Println(recover())	//If call recover in a program that isn't panicking, it does nothing and returns nil.
	freakOut(panicValue)
	fmt.Println(exitValue)
}
