package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// print Go executable path in your system.
func main() {
	outputVersionPath()
	cmdStartWait()
}

func outputVersionPath() {
	goExecPath, err := exec.LookPath("go")
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Go Executable: ", goExecPath)
	}

	// created a basic Cmd struct which points to standard Go executable path.
	// The Args field contains the list of arguments to invoke the Go executable.
	// set STDOUT and STDERR to os.Stdout that will spit out the result of the execution to the terminal.
	cmdGoVersion := &exec.Cmd{
		Path:   goExecPath,
		Args:   []string{goExecPath, "version"},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	// Print cmdGoVersion struct, but not run
	// fmt.Println(cmdGoVersion.String())

	// run the cmdGoVersion, it is the Cmd.Run method runs the command and throws an error if the command could not successfully run.
	// This error does not contain standard output.
	if err1 := cmdGoVersion.Run(); err1 != nil {
		fmt.Println("Error:", err) // the standard output is: go version go1.18 darwin/arm64
	}
}

func cmdStartWait() {
	cmd := exec.Command("sleep", "3")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	_ = cmd.Wait()
	fmt.Println("Hello CMD")
}
