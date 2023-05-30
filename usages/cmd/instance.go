package main

import (
	"fmt"
	"log"
	"os/exec"
)

// print Go executable path in your system.
func main() {
	outputVersionPath()
	cmdStartWait()
	inputCmd()
}

func outputVersionPath() {
	cmd := exec.Command("ls", "-lah")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The output is: %s\n", out)
}

func cmdStartWait() {
	cmd := exec.Command("sleep", "3")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	_ = cmd.Wait()
	log.Println("Hello CMD")
}

func inputCmd() {
	// Prompt the user to enter a number
	fmt.Print("Enter a number: ")

	var input float64
	a := &input

	// Scanf fills input with the number we enter
	_, err := fmt.Scanf("%f", a)
	if err != nil {
		return
	}
	output := input * 2
	fmt.Println(output)
}
