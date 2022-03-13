package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func scanDir(path string) error {
	fmt.Println(path) // print the current directory.
	// Get a slice with the directory's contents
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range files {
		// Join the directory path and filename with a slash.
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() { // If this file is a subdirectory
			err := scanDir(filePath) // recursively call scanDir(), this time with the subdirectory's path.
			if err != nil {
				return err
			}
			fmt.Println("Directory:", file.Name())
		} else { // Otherwise, print File and the filename.
			fmt.Println(filePath) // If it is a regular file, just print its path.
		}
	}
	return nil
}

/**
Using the below command to run this file.
go run base/usages/files/files.go base/usages
go run base/usages/files/files.go base
*/
func main() {
	err := scanDir(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
}
