package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// []string means return a slice of strings
func readFile() ([]string, error) {
	// Getwd() returns a rooted path name corresponding to the current directory
	p, _ := os.Getwd()
	fmt.Println(p)

	/*
		In Go, the readFile path should use the absolute path, if not, when run it, the path should the same with
		current run path. Otherwise, it won't find the file.
		os.Open() is used for open a file in go.

		The file path can become a variable, then it will easy to read more files.
	*/
	file, err := os.Open("./practices/votes.txt")
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
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()

		// Append a line to the slice.
		lines = append(lines, line)
	}

	// fmt.Println(lines)
	return lines, nil
}

func main() {
	// The readFile method can be improved as a package on github.
	lines, err := readFile()

	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(lines)

	// Declare a map that will use candidate names as keys, and vote counts as values.
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}

	for name, count := range counts {
		fmt.Printf("Votes for %s: %d\n", name, count)
	}
}
