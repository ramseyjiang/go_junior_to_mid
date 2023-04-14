package opreflect

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func UseReflectStructFields(s interface{}) {
	val := reflect.ValueOf(s)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		_ = val.Type().Field(i).Name
		_ = field.Interface()
		// fmt.Printf("%v: %v\n", val.Type().Field(i).Name, field.Interface())
	}
}

func UnusedStructFields(p Person) {
	_ = p.Name
	_ = p.Age
	// fmt.Printf("Name: %v\nAge: %v\n", p.Name, p.Age)
}

func BenchmarkReflect(b *testing.B) {
	p := Person{"John Doe", 30}
	for i := 0; i < b.N; i++ {
		UseReflectStructFields(p)
	}
}

func BenchmarkNonReflect(b *testing.B) {
	p := Person{"John Doe", 30}
	for i := 0; i < b.N; i++ {
		UnusedStructFields(p)
	}
}
