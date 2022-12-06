package main

import (
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
	// LookPath can be used to check whether the cmd exists or not.
	checkCmd, err := exec.LookPath("la")
	if err != nil { // it will output "Error:  exec: "la": executable file not found in $PATH"
		log.Println("Error: ", err)
	} else {
		log.Println("Check la path is ", checkCmd)
	}

	goExecPath, err := exec.LookPath("go")
	if err != nil {
		log.Println("Error: ", err)
	} else {
		log.Println("Go Executable: ", goExecPath)
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

	nameEnv := "NAME=golang"
	ageEnv := "AGE=13"
	newEnv := append(os.Environ(), nameEnv, ageEnv)
	cmdGoVersion.Env = newEnv

	// Print cmdGoVersion struct, but not run
	// fmt.Println(cmdGoVersion.String())

	// run the cmdGoVersion, it is the Cmd.Run method runs the command and throws an error if the command could not successfully run.
	// This error does not contain standard output.
	if err1 := cmdGoVersion.Run(); err1 != nil {
		log.Println("Error:", err) // the standard output is: go version go1.18 darwin/arm64
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
	log.Println("Hello CMD")
}
