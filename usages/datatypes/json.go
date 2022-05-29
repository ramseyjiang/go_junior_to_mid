package main

import (
	"encoding/json"
	"log"
)

type response struct {
	PageNumber int      `json:"page,omitempty"`
	Fruits     []string `json:"fruits"`
	FirstName  string   `json:"first_name,omitempty"`
	LastName   string   `json:"-"`
}

type request struct {
	Apple    int16  `json:"apple"`
	Lettuce  int16  `json:"lettuce"`
	LastName string `json:"LastName,omitempty"`
}

func encodeUsage() {
	mapA := map[string]int{"apple": 5, "lettuce": 7}
	mapB, err := json.Marshal(&mapA) // Marshal return is []byte and error. It means a slice and error.
	if err != nil {
		log.Println(err)
	}
	log.Printf("Json encode1 usage is as below: %v", string(mapB))

	mapC := request{}
	mapC.Apple = 5
	mapC.Lettuce = 7
	mapD, err := json.Marshal(&mapC)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Json encode2 usage is as below: %v", string(mapD))
}

func decodeUsage(str string) {
	res := response{}

	// While decoding the json byte using unmarshal.
	// the first argument is the json byte slice, function unmarshal only accept json data type []byte
	// the second argument is the address of the response type struct where we want the json to be mapped to
	if err := json.Unmarshal([]byte(str), &res); err != nil {
		panic(err)
	}

	log.Printf("Json decode usage is as below: %v", res)
	log.Printf("Page is %v, fruits are %v\n", res.PageNumber, res.Fruits)
}

func main() {
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	encodeUsage()
	decodeUsage(str)

	// the last_name won't be show after Unmarshal, because "-" in json, it means ignore.
	receiver1 := `{"page": 0, "fruits": ["apple", "peach"], "first_name": "Ramsey", "last_name": "xxx"}`
	receiver2 := `{"page": 2, "fruits": ["apple", "peach"], "first_name": "", "last_name": "xxx"}`

	// Output is {0 [apple peach] Ramsey}, because even pag is 0, it is still an int number, it cannot output empty.
	decodeUsage(receiver1)

	// Output is {2 [apple peach] }.
	// first_name is empty "" in the json, so after Unmarshal, it cannot show in the output.
	// That's the difference from int and string during using json.
	decodeUsage(receiver2)
}
