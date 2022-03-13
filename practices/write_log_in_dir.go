package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

var writingSync sync.Mutex

func main() {
	logDirectoryCheck()
	logInfo("run", "Program is running")
}

func logInfo(reference, data string) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000") + " [" + reference + "] --INF-- " + data)
	appendDataToLog("INF", reference, data)
}

func logDirectoryCheck() {
	dateTimeFormat := "2006-01-02 15:04:05.000"
	dir := checkEnvType(dateTimeFormat)
	createLogDirectory(dir, dateTimeFormat)
}

func createLogDirectory(dir string, dateTimeFormat string) {
	logDirectory := filepath.Join(dir, "logs")
	_, checkPathError := os.Stat(logDirectory)
	logDirectoryExists := checkPathError == nil
	if logDirectoryExists {
		return
	}
	switch runtime.GOOS {
	case "windows":
		{
			err := os.Mkdir(logDirectory, 0777)
			if err != nil {
				fmt.Println(time.Now().Format(dateTimeFormat) + " [MAIN] --ERR-- Unable to create directory for log file: " + err.Error())
				return
			}
			fmt.Println(time.Now().Format(dateTimeFormat) + " [MAIN] --INF-- Log directory created")
		}

	default:
		{
			err := os.MkdirAll(logDirectory, 0777)
			if err != nil {
				fmt.Println(time.Now().Format(dateTimeFormat) + " [MAIN] --ERR-- Unable to create directory for log file: " + err.Error())
				return
			}
			fmt.Println(time.Now().Format(dateTimeFormat) + " [MAIN] --INF-- Log directory created")
		}
	}
}

func checkEnvType(dateTimeFormat string) string {
	var dir string
	switch runtime.GOOS {
	case "windows":
		{
			executable, err := os.Executable()
			if err != nil {
				fmt.Println(time.Now().Format(dateTimeFormat) + " [MAIN] --ERR-- Unable to read actual directory: " + err.Error())
			}
			dir = filepath.Dir(executable)
		}
	default:
		{
			executable, err := os.Getwd()
			if err != nil {
				fmt.Println(time.Now().Format(dateTimeFormat) + " [MAIN] --ERR-- Unable to read actual directory: " + err.Error())
			}
			dir = executable
		}
	}
	return dir
}

func appendDataToLog(logLevel string, reference string, data string) {
	dateTimeFormat := "2006-01-02 15:04:05.000"
	logNameDateTimeFormat := "2006-01-02"
	logDirectory := filepath.Join(".", "logs")
	logFileName := reference + " " + time.Now().Format(logNameDateTimeFormat) + ".log"
	logFullPath := strings.Join([]string{logDirectory, logFileName}, "/")
	logData := time.Now().Format("2006-01-02 15:04:05.000 ") + reference + " " + logLevel + " " + data
	writingSync.Lock()
	f, err := os.OpenFile(logFullPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(time.Now().Format(dateTimeFormat) + " [" + reference + "] --ERR-- Cannot open file: " + err.Error())
		writingSync.Unlock()
		return
	}
	_, err = f.WriteString(logData + "\r\n")
	if err != nil {
		fmt.Println(time.Now().Format(dateTimeFormat) + " [" + reference + "] --ERR-- Cannot write to file: " + err.Error())
	}
	err = f.Close()
	if err != nil {
		fmt.Println(time.Now().Format(dateTimeFormat) + " [" + reference + "] --ERR-- Cannot close file: " + err.Error())
	}
	writingSync.Unlock()
}
