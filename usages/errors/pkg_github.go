package main

import (
	"fmt"
	"github.com/pkg/errors"
)

// Go does not provide the capability to a stack trace to an error, that's why I use the pkg "github.com/pkg/errors"
func main() {
	originalErr := errors.New("This is original error message.")

	// errors.Wrap() provides a stack trace which includes exact information about where the error occurred or returned in our code.
	// When an error occurs, a stack trace is a great way to debug your code as it contains the filename and the line number
	// at which the error has occurred and a stack of function calls made until the error occurred.
	err := errors.Wrap(originalErr, "Context info using Wrap from errors github package.\n")

	fmt.Printf("Usual => %v \n\n", err)

	fmt.Printf("with stack trace => %+v \n\n", err)

	// Any error which implements causer interface which contains Cause() error method can be inspected by the Cause function.
	originalErrUnWrapped := errors.Cause(err) // Cause method which used to extract original error.
	fmt.Println(originalErrUnWrapped)
}
