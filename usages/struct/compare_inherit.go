package main

import (
	"fmt"
)

// The whole file has several examples about parents and children classes as other languages.

// The struct in Go, it is similar as other languages define a class.
// The address, in the subscriber and employee, belongs to an embedded struct.
// You can access the fields of an embedded struct as if they belong to the outer struct.
// Define a type named "subscriber".
// Defined types also work for function parameters and return values.
// If the type subscriber use outside this package, must remember each name exported should be capitalized.
type subscriber struct {
	name        string
	number      float64
	active      bool
	homeAddress address // A struct can be a field on another struct as well. This way is the explicit way.
}

type employee struct {
	name    string
	salary  float64
	address // This way is the anonymous struct fields.
}

type address struct {
	city     string
	street   string
	postCode int
}

// In struct, it does not have any restrict with data type. In array and slice, both of them have only one type data.
func printStructDefaultValue() {
	var myStruct struct {
		number float64
		word   string
		toggle bool
		index  int
	}

	// It will print out all type in the struct default value
	// When we call Printf with the %#v verb above, it prints the value in myStruct as a struct literal.
	fmt.Printf("%#v\n", myStruct)
	// Print result is below
	// struct { number float64; word string; toggle bool; index int }{number:0, word:"", toggle:false, index:0}

	// Access struct fields using the dot operator
	fmt.Printf("myStruct.index is %v, myStruct.word is %v, myStruct.toggle is %v, myStruct.number is %v\n",
		myStruct.index, myStruct.word, myStruct.toggle, myStruct.number) // 0 "" false 0
}

// Use defined type as the variable types.
// It includes two different ways to define a struct and then assign values.
func setSubscriberValue() {
	// The first way, declare a variable of type "subscriber".
	var subscriber1, subscriber2 subscriber

	subscriber1.name = "Alice"
	subscriber1.number = 13
	subscriber1.active = true
	subscriber2.name = "Amy"
	// subsciber.number = "test", it will have an error. Because it is not the int type.

	fmt.Println(showSubscriberInfo(subscriber1))
	fmt.Println(showSubscriberInfo(subscriber2))

	// The second way, define a struct and then assign values
	secSubscriber := subscriber{name: "may", active: true, number: 10}
	fmt.Println(secSubscriber)
}

// The function name is a way to use defined type as a parameter.
// user subscriber, the user is declared the subscriber type.
// The return value is another way to use defined type as a return value. The return is the subscriber type.
// The subscriber after the "(user subscriber)", it declares the return value should be the subscriber type.
func showSubscriberInfo(user subscriber) subscriber {
	// fmt.Println(user)
	return user
}

// It is an easy way to change a value in struct and don't need a return.
// Using this way, it will take up computer's memory. So if lots of fields, using this way not a good choice.
func applyDiscount(testSubscriber *subscriber) {
	// Assign to the struct field through the pointer.
	testSubscriber.number = 5
}

func setAddressValue(addressStruct *address) {
	addressStruct.city = "Auckland"
	addressStruct.street = "Wattle downs"
	addressStruct.postCode = 2103
}

func setEmployeeValue(employeeStruct *employee) {
	employeeStruct.name = "Davy"
	employeeStruct.salary = 12000
}

func main() {
	printStructDefaultValue()
	setSubscriberValue()

	var testSubscriber subscriber
	// Using the pointer way to change the struct. In the applyDiscount func, it does not need return.
	applyDiscount(&testSubscriber)
	fmt.Println(testSubscriber)

	var addressStruct address
	setAddressValue(&addressStruct)
	fmt.Println(addressStruct)

	// If it is the explicit, it uses the name directly.
	testSubscriber.homeAddress = addressStruct
	fmt.Println(testSubscriber, testSubscriber.homeAddress.city)

	var employeeStruct employee
	setEmployeeValue(&employeeStruct)

	// If it is the anonymous fields, it uses an anonymous field to make our inner struct easier to access.
	employeeStruct.address = addressStruct
	fmt.Println(employeeStruct, employeeStruct.address.postCode)
}
