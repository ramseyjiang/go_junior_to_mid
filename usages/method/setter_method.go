package main

import (
	"errors"
	"fmt"
	"log"
)

// Setter methods are methods used to set fields or other values within a defined type’s underlying value.
// Go still uses a Set prefix for setter methods. because it’s needed to distinguish setter method names from getter
// method names for the same field.
// The Go community has decided on a convention of leaving the Get prefix off of getter method names.
// Including it would only lead to confusion for your fellow developers!

type Date struct {
	year  int
	month int
	day   int
}

type Event struct {
	title string
	Date
}

// To define a method, provide a receiver parameter in parentheses before the method name.
// (d *Date) is the receiver parameter.
// d is the receiver parameter value.
// *Date is the receiver parameter type. It is the value to call method on.
// The *Date is a pointer receiver, so original value can be updated.

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	} else {
		// d is the receiver parameter value, so it will automatically get value at pointer.
		d.year = year
		return nil
	}
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	} else {
		// d is the receiver parameter value, so it will automatically get value at pointer.
		d.month = month
		return nil
	}
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	} else {
		// d is the receiver parameter value, so it will automatically get value at pointer.
		d.day = day
		return nil
	}
}

func (e *Event) SetTitle(title string) error {
	// utf8.RuneCountInString(title) also can be used to calculate the length of title.
	if len(title) > 30 {
		return errors.New("invalid title")
	}
	e.title = title
	return nil
}

// Common base error process
func processErr(err error, nil error) {
	if err != nil {
		log.Fatal(err)
	}
}

func SetEvent(event *Event) {
	// err := event.SetTitle("An extremely long title that is impractical to print")	//It will have error cause title over 30.
	err := event.SetTitle("An title is not long")
	processErr(err, nil)

	err = event.SetYear(2021)
	processErr(err, nil)

	err = event.SetMonth(11)
	processErr(err, nil)

	err = event.SetDay(16)
	processErr(err, nil)
}

func main() {
	event := Event{}
	SetEvent(&event)

	fmt.Println(event, event.title, event.year)
}
