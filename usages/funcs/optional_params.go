package main

import (
	"log"
)

type Student struct {
	// required fields
	Id   int
	Name string
	Sex  bool

	// optional
	Country string
}

type Option func(f *Student)

func WithCounty(country string) Option {
	return func(f *Student) {
		f.Country = country
	}
}

func NewStudent(id int, name string, sex bool, opts ...Option) *Student {
	student := &Student{
		Id:      id,
		Name:    name,
		Sex:     sex,
		Country: "China", // default value
	}

	for _, opt := range opts {
		opt(student)
	}

	return student
}

func main() {
	s1 := NewStudent(1, "Jerry", true)
	log.Println(s1)

	s2 := NewStudent(2, "Tom", true, WithCounty("USA"))
	log.Println(s2)
}
