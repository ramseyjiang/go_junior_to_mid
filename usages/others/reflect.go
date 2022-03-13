package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string
	age  int64
}

// Using TypeOf and ValueOf gives back to us the underlying type, and a pointer to the value.
// The value coming from ValueOf is of type Reflect.Value, and not the original variable type.
// Reflect.Kind is like the types of types. Every variable has a kind, derived from its type.
func main() {
	type1, value1, valueInt64 := emptyInterfaceReflect()
	fmt.Println(type1, value1, valueInt64+1) // output is int 10 11

	// Though value1 is 10, it cannot do the below, because it is the reflection value,
	// Go compiler does not recognize this action for Reflect.Value. but just for Integer types.
	// fmt.Println(value1 + 1) // It will have an error as the comments above.

	type2, value2 := structReflect()
	fmt.Println(type2, value2) // output is main.Person {Kobe Davy}

	var myVar interface{} = Person{name: "Ramsey", age: 40}
	getTypeFromStruct(myVar)
}

func emptyInterfaceReflect() (reflect.Type, reflect.Value, int64) {
	var myVar interface{} = 10

	type1 := reflect.TypeOf(myVar)
	value1 := reflect.ValueOf(myVar)
	intValue1 := value1.Int()
	return type1, value1, intValue1
}

func structReflect() (reflect.Type, reflect.Value) {
	var person interface{} = Person{name: "Kobe Davy"}

	type2 := reflect.TypeOf(person)
	value2 := reflect.ValueOf(person)
	return type2, value2
}

func getTypeFromStruct(s interface{}) {
	reflectValue := reflect.ValueOf(s)

	// Make sure we are handling with a struct here
	if reflectValue.Kind() == reflect.Struct {
		fieldCount := reflectValue.NumField()
		fmt.Println("Num of fields: ", fieldCount)

		for i := 0; i < fieldCount; i++ {
			// Get individual field details
			field := reflectValue.Field(i)
			fmt.Printf("type: %T, value: %v \n", field, field)

			// It works, because Person.age can be changed to string type from int type.
			fmt.Printf("type: %T, value: %v \n", field.String(), field)

			// The below won't work, because Person.name is a string type.
			// fmt.Printf("type: %T, value: %v \n", field.Int(), field)
		}
	}
}
