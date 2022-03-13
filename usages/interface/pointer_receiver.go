package main

import "fmt"

type Switch string

// The toggle method on the Switch type below has to use a pointer receiver, so it can modify the receiver.
func (s *Switch) toggle() {
	if *s == "on" {
		*s = "off"
	} else {
		*s = "on"
	}
	fmt.Println(*s)
}

type Toggleable interface {
	toggle()
}

// If a type declares methods with pointer receivers, then you’ll only be able to use pointers to that type when
// assigning to interface variables.
func main() {
	s := Switch("off")

	// when we assign a Switch value to a variable with the interface type Toggleable, should assign a pointer to it.
	var t Toggleable = &s

	t.toggle()
	t.toggle()
}
