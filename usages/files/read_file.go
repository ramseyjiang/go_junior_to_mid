package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Getwd() returns a rooted path name corresponding to the current directory
	p, _ := os.Getwd()
	println(p)

	/*
		In Go, the readFile path should use the absolute path, if not, when run it, the path should the same with
		current run path. Otherwise, it won't find the file.
		os.Open() is used for open a file in go.

		The file path can become a variable, then it will easy to read more files.
	*/
	file, err := os.Open("./base/usages/files/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	/*
		bufio.NewScanner() is used to read a file in go.
		scanner.Scan() is used as a part of a for loop. It will read a single line of text from the file, returning
		true if it read data successfully and false if it did not.
		If Scan is used as the condition on a for loop, the loop will continue running as long as there is more
		data to be read. Once the end of the file is reached (or there’s an error), Scan will return false, and
		the loop will exit.
	*/
	var numbers []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())

		// Convert the string to a float64 and assign it to a temporary variable.
		number, err := strconv.ParseFloat(scanner.Text(), 64)

		if err != nil {
			log.Fatal(err)
		}

		// Append a number to the slice.
		numbers = append(numbers, number)
	}

	/*
		Keeping files open consumes resources from the operating system, so files should always be closed when
		a program is done with them. Calling the Close method on the os.File will accomplish this.
	*/
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	/*
		It’s also possible that the bufio.Scanner encountered an error while scanning through the file.
		If it did, calling the Err method on the scanner will return that error, which we log before exiting.
	*/
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	fmt.Println(numbers)
}
