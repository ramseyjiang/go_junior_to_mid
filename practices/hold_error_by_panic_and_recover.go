package main

import "fmt"

func calmDown() {
	// If recover is called when there is no panic, it simply returns nil.
	// If recover is called during a panic, it returns the original value that was passed to panic function.
	p := recover()

	// This if condition is used to fix the ISSUE which other panic happens.
	if p == nil {
		return
	}

	// Assert that the type of the panic value is "error" .
	// error is an interface, after it is asserted, err can use err.Error() method directly.
	err, ok := p.(error)

	if ok {
		fmt.Println(err.Error()) // It has an "error" value, after it calls the Error method.
	} else {
		panic(p) // If the panic value is not an error, resume panicking with the original value.
	}
}

// To call methods or do anything else with the panic value, you’ll need to convert it back to its underlying type
// using a type assertion.
// The below code that takes the return value of recover() and convert it back to an error value.
// The output is: there is an error
func main() {
	errContent := "these is an error"
	defer calmDown()
	// panic("some other issue")	//Here, if uncomment this line, the code will stop and won't print any panic. ISSUE.
	err := fmt.Errorf(errContent)

	// Any value can be passed into panic will be converted to a string and printed as part of the log message.
	// A panic log message includes a stack trace, a list of all active function calls that can be useful for debugging.
	panic(err)
}
