package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	// fmt.Errorf inserts values into a format string just like fmt.Printf or fmt.Sprintf, but instead of printing or returning a string, it returns an error value.
	err2 := fmt.Errorf("A height of %0.2f is invalid", -2.56)

	fmt.Println(err2.Error()) // One way outputs an error.
	fmt.Println(err2)         // The other way outputs an error.

	// Pass a string to the errors package’s New function, which will return a new error value.
	err1 := errors.New("height can't be negative")
	fmt.Println(err1) // call the Error method on that error value, you’ll get the string which is passed to errors.New.
	log.Fatal(err1)
}
