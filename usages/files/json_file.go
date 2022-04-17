package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type car struct {
	Speed int    `json:"speed"`
	Brand string `json:"brand"`
}

const fileName = "./jsonFile.json"

func writeJsonFile(fileName string) error {
	c := car{
		Speed: 10,
		Brand: "Nio",
	}
	dat, err := json.Marshal(c)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fileName, dat, 0644)
	if err != nil {
		return err
	}
	fmt.Println("Write JSON file has been done.")

	return nil
}

func readJsonFile(fileName string) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	c := car{}
	err = json.Unmarshal(data, &c)
	if err != nil {
		return err
	}
	fmt.Println("JSON file content is ", string(data))

	return nil
}

func main() {
	_ = writeJsonFile(fileName)
	_ = readJsonFile(fileName)
}
