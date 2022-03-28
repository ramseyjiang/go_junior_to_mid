package main

import (
	"encoding/json"
	"fmt"
)

type response struct {
	PageNumber int      `json:"page"`
	Fruits     []string `json:"fruits"`
}

func encodeUsage() {
	mapA := map[string]int{"apple": 5, "lettuce": 7}
	mapB, err := json.Marshal(mapA)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Json encode usage is as below:")
	fmt.Println(string(mapB))
}

func decodeUsage(str string) {
	res := response{}

	// While decoding the json byte using unmarshal.
	// the first argument is the json byte, function unmarshal only accept json data type []byte
	// the second argument is the address of the response type struct where we want the json to be mapped to
	if err := json.Unmarshal([]byte(str), &res); err != nil {
		panic(err)
	}

	fmt.Println("Json decode usage is as below:")
	fmt.Println(res) // res type is defined in the struct at the top.
	fmt.Printf("Page is %v, fruits are %v", res.PageNumber, res.Fruits)
}

func main() {
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	encodeUsage()
	decodeUsage(str)
}
