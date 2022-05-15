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

func encodeUsage() {
	mapA := map[string]int{"apple": 5, "lettuce": 7}
	mapB, err := json.Marshal(mapA) // Marshal return is []byte and error. It means a slice and error. Click Marshal to see.
	if err != nil {
		log.Println(err)
	}
	log.Println("Json encode usage is as below:")
	log.Println(string(mapB))
}

func decodeUsage(str string) {
	res := response{}

	// While decoding the json byte using unmarshal.
	// the first argument is the json byte slice, function unmarshal only accept json data type []byte
	// the second argument is the address of the response type struct where we want the json to be mapped to
	if err := json.Unmarshal([]byte(str), &res); err != nil {
		panic(err)
	}

	log.Println("Json decode usage is as below:")
	log.Println(res) // res type is defined in the struct at the top.
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

	// jsonTypes()
}

// type ServerConf struct {
// 	Host         string        `json:",default=0.0.0.0"`
// 	Port         int           `json:",range=[80,65535)"`
// 	LogMode      string        `json:",options=[file,console]"`
// 	Verbose      bool          `json:",optional"`
// 	MaxConn      int           `json:",default=10000"`
// 	Timeout      time.Duration `json:",default=3s"`
// 	CpuThreshold int64         `json:",default=900,range=[0:1000]"`
// }
//
// func jsonTypes() {
// 	serverConf := ServerConf{}
// 	log.Printf("Default Host is %v\n", serverConf.Host)
// 	log.Printf("Default Port is %v\n", serverConf.Port)
// 	log.Printf("Default LogMode is %v\n", serverConf.LogMode)
// 	log.Printf("Default Verbose is %v\n", serverConf.Verbose)
// 	log.Printf("Default MaxConn is %v\n", serverConf.MaxConn)
// 	log.Printf("Default Timeout is %v\n", serverConf.Timeout)
// 	log.Printf("Default CpuThreshold is %v\n", serverConf.CpuThreshold)
//
// 	server := `{"Host": "127.0.0.1", "CpuThreshold": -10}`
// 	decodeServerJson(server)
// }
//
// func decodeServerJson(str string) {
// 	server := ServerConf{}
//
// 	if err := json.Unmarshal([]byte(str), &server); err != nil {
// 		panic(err)
// 	}
//
// 	log.Println("Json decode usage is as below:")
// 	log.Println(server)
// }
