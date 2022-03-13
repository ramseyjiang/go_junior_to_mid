package main

import (
	"fmt"
	"reflect"
	"time"
)

/*
The time.Now function returns a new Time value for the current date and
time, which we store in the now variable.
*/
func main() {
	now := time.Now()
	fmt.Println(now)

	var year = now.Year()
	var month = int(now.Month()) // Here if use now.Month() directly, it will have a type error. So it uses int(now.Month())
	var day = now.Day()
	var hour = now.Hour()
	var minute = now.Minute()
	var second = now.Second()

	fmt.Println(reflect.TypeOf(now.Month()), now.Month()) // time.Month November
	fmt.Println(year, month, day, hour, minute, second)
}
