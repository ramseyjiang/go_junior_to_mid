package main

import (
	"encoding/json"
	"fmt"
)

type Client struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	var jsonString = `{"Name": "john doe", "Age": 27}`
	var jsonData = []byte(jsonString)
	var data Client

	// Function unmarshal only accept json data type []byte
	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// FullName field has tag json:"Name" .
	// This tag is used for mapping json information to the related field .
	// Parsed Json data have 2 field which Name and Age .
	// Fortunately, field Age has the same name as json data.
	// If field Age does not exist, if it was a count field, then data.count will output 0.
	fmt.Println(data)
	fmt.Println("user :", data.FullName)
	fmt.Println("age  :", data.Age)
}
