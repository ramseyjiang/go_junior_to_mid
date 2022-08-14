package main

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type car struct {
	Speed int    `json:"speed"`
	Brand string `json:"brand"`
	Color string `json:"color"`
}

const fileName0 = "unArr.json"
const fileName1 = "arr.json"
const fileName2 = "arr.csv"

func writeJsonFile(fileName string) error {
	c := car{
		Speed: 10,
		Brand: "Nio",
		Color: "Blue",
	}
	dat, err := json.Marshal(c)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fileName, dat, 0644)
	if err != nil {
		return err
	}
	log.Println("Write JSON file has been done.")

	return nil
}

func readJsonFile(fileName string) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	c := car{} // Different with read array json file
	err = json.Unmarshal(data, &c)
	if err != nil {
		return err
	}
	log.Println("JSON file content is ", string(data))

	return nil
}

func writeArrJsonFile(fileName string) error {
	nio := car{Speed: 60, Brand: "Nio", Color: "Silver"}
	li := car{Speed: 50, Brand: "Li", Color: "Red"}
	tesla := car{Speed: 80, Brand: "tesla", Color: "Black"}
	cars := []car{nio, li, tesla}
	data, err := json.Marshal(cars)
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(fileName, data, 0644); err != nil {
		return err
	}
	log.Println("Write Array JSON file has been done.")

	return nil
}

func readArrJsonFile(fileName string) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	var cars []car // Different with read json file
	err = json.Unmarshal(data, &cars)
	if err != nil {
		return err
	}
	log.Println("JSON file content is ", cars)

	return nil
}

func openClose(fn string) {
	f, err := os.Create(fn)
	if err != nil {
		log.Println(err)
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			log.Println(err)
		}
	}(f)
}

func writeCSVFile(fileName string) {
	f, err := os.Create(fileName) // If read a csv, here should use os.Open()
	if err != nil {
		log.Println(err)
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			log.Println(err)
		}
	}(f)

	w := csv.NewWriter(f)
	records := [][]string{
		{"id", "fruit", "color", "taste"},
		{"0", "apple", "red", "sweet"},
		{"1", "banana", "yellow", "sweet"},
		{"2", "lemon", "yellow", "sour"},
		{"3", "grapefruit", "red", "sour"},
	}
	_ = w.WriteAll(records)
}

func readCSVFile(fileName string) {
	f, err := os.Open(fileName) // If write a csv, here should use os.Create()
	if err != nil {
		log.Println(err)
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			log.Println(err)
		}
	}(f)

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Println(err)
	}
	log.Println(records)
}

func main() {
	openClose(fileName0)
	_ = writeJsonFile(fileName0)
	_ = readJsonFile(fileName0)

	openClose(fileName1)
	_ = writeArrJsonFile(fileName1)
	_ = readArrJsonFile(fileName1)

	writeCSVFile(fileName2)
	readCSVFile(fileName2)
}
