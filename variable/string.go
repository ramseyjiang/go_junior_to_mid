package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {
	base()

	str0 := withoutStringsBuilderWithFmt()
	log.Println(str0)

	str1 := withoutStringsBuilder()
	log.Println(str1)

	str2 := stringsBuilderWithForRange()
	log.Println(str2)

	str3 := stringBuilderWithForLoop()
	log.Println(str3)
}

func withoutStringsBuilderWithFmt() (allStr string) {
	execTime := time.Now()
	s := [5]string{"one", "two", "3", "4", "five"}
	for _, str := range s { // change to for loop
		allStr += str
	}
	fmt.Printf("withoutStringsBuilderWithFmt ExecTime is %33s %6s\n", "->", time.Since(execTime))
	return
}

func withoutStringsBuilder() (allStr string) {
	execTime := time.Now()
	s := [5]string{"one", "two", "3", "4", "five"}
	for _, str := range s { // change to for loop
		allStr += str
	}
	log.Printf("withoutStringsBuilder ExecTime is %20s %6s\n", "->", time.Since(execTime))
	return
}

// If not declare result at the top as str string, it will slow around 30 ns for the whole func run time.
func stringBuilderWithForLoop() (str string) {
	var sb strings.Builder
	s := [5]string{"one", "two", "3", "4", "five"}

	execTime := time.Now()
	for i := 0; i < len(s); i++ {
		sb.WriteString(s[i])
	}
	log.Printf("stringBuilderWithForLoop ExecTime is %17s %6s\n", "->", time.Since(execTime))

	log.Printf("Last %28s Len : %3d, Cap : %3d, ptr : %3p\n", " ", sb.Len(), sb.Cap(), &sb)

	return sb.String()
}

func stringsBuilderWithForRange() (str string) {
	var sb strings.Builder
	s := [5]string{"one", "two", "3", "4", "five"}
	// fmt.Printf("Initial %25s Len : %3d, Cap : %3d, ptr : %3p\n", " ", sb.Len(), sb.Cap(), &sb)

	// var str string  //If not declare str result at the top, instead declare at here, it will slow at least 0.2 µs for whole func.
	execTime := time.Now()
	for _, str = range s {
		sb.WriteString(str)
		// fmt.Printf("String : %5s, String Len : %3d, Len : %3d, Cap : %3d, ptr : %3p\n", str, len(str), sb.Len(), sb.Cap(), &sb)
	}
	log.Printf("stringsBuilderWithForRange ExecTime is %15s %6s\n", "->", time.Since(execTime))

	// str = sb.String()
	sb.Reset()
	log.Printf("After reset %21s Len : %3d, Cap : %3d, ptr : %3p\n", " ", sb.Len(), sb.Cap(), &sb)
	return sb.String()
}

func base() {
	s := "this is a string"
	log.Printf("%v, %T\n", s, s) // this is a string, string

	log.Printf("%v, %T\n", s[2], s[2]) // 105, uint8. Because it should do the type convert first

	log.Printf("%v, %T\n", string(s[2]), string(s[2])) // i, string

	b := []byte(s)
	log.Printf("%v, %T\n", b, b) // [116 104 105 115 32 105 115 32 97 32 115 116 114 105 110 103], []uint8
}
