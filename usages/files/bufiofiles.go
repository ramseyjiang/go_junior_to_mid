package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

/**
usages/files/validdata.txt is the reading file name.
% go run usages/files/bufiofiles.go usages/files/validdata.txt

/Users/daweijiang/Desktop/GoApp/go_junior_to_mid
Opening usages/files/validdata.txt
Closing file
Sum: 252.80

usages/files/invaliddata.txt is the reading file name, if file content are incorrect, it will have an error.
% go run usages/files/bufiofiles.go usages/files/invaliddata.txt

2022/08/14 14:46:46 Current file path is: /Users/daweijiang/Desktop/GoApp/go_junior_to_mid
Opening file name is usages/files/invaliddata.txt
Closing file
2022/08/14 14:46:46 strconv.ParseFloat: parsing "hello": invalid syntax
exit status 1


The file is closed by both of them, because the incorrect one is used defer to trigger it.
*/
func main() {
	// Getwd() returns a rooted path name corresponding to the current directory
	p, _ := os.Getwd()
	log.Println("Current file path is:", p)

	numbers, err := GetFloats(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	log.Printf("Sum: %0.2f\n", sum)
}

func OpenFile(fileName string) (*os.File, error) {
	log.Println("Opening file name is", fileName)
	// os.Open() is used for open a file in go.
	return os.Open(fileName)
}

func CloseFile(file *os.File) {
	log.Println("Closing file")
	_ = file.Close()
}

func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := OpenFile(fileName)
	if err != nil {
		return nil, err
	}

	defer CloseFile(file) // Close the file even if it has an error.

	/*
		bufio.NewScanner() is used to read a file in go.
		scanner.Scan() is used as a part of a for loop.
		It will read a single line of text from the file, returning true if it read data successfully and false if it did not.
		If Scan is used as the condition on a for loop, the loop will continue running as long as there is more data to be read.
		Once the end of the file is reached (or there’s an error), Scan will return false, and the loop will exit.
	*/
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Convert the string to a float64 and assign it to a temporary variable.
		number, err := strconv.ParseFloat(scanner.Text(), 64) // convert string to float64
		if err != nil {
			return nil, err
		}
		log.Println("file content includes number:", number)
		numbers = append(numbers, number)
	}

	// It’s possible that the bufio.Scanner encountered an error while scanning through the file.
	// If it did, calling the Err method on the scanner will return that error, which we log before exiting.
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}
