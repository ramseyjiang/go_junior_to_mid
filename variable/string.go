package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	base()
	str1 := withoutStringsBuilder()
	fmt.Println(str1)

	str2 := stringsBuilderWithForRange()
	fmt.Println(str2)

	str3 := stringBuilderWithForLoop()
	fmt.Println(str3)
}

func withoutStringsBuilder() (allStr string) {
	execTime := time.Now()
	s := [5]string{"one", "two", "3", "4", "five"}
	for _, str := range s { // change to for loop
		allStr += str
	}
	fmt.Printf("ExecTime is %20s %6s\n", "->", time.Since(execTime))
	return
}

func stringBuilderWithForLoop() (str string) {
	var sb strings.Builder
	s := [5]string{"one", "two", "3", "4", "five"}

	execTime := time.Now()
	for i := 0; i < len(s); i++ {
		sb.WriteString(s[i])
	}
	fmt.Printf("ExecTime is %20s %6s\n", "->", time.Since(execTime))

	fmt.Printf("Last %28s Len : %3d, Cap : %3d, ptr : %3p\n", " ", sb.Len(), sb.Cap(), &sb)

	return sb.String()
}

func stringsBuilderWithForRange() (str string) {
	var sb strings.Builder
	s := [5]string{"one", "two", "3", "4", "five"}
	fmt.Printf("Initial %25s Len : %3d, Cap : %3d, ptr : %3p\n", " ", sb.Len(), sb.Cap(), &sb)

	execTime := time.Now()
	for _, str = range s {
		sb.WriteString(str)
		fmt.Printf("String : %5s, String Len : %3d, Len : %3d, Cap : %3d, ptr : %3p\n", str, len(str), sb.Len(), sb.Cap(), &sb)
	}
	fmt.Printf("ExecTime is %20s %6s\n", "->", time.Since(execTime))

	// str = sb.String()
	sb.Reset()
	fmt.Printf("After reset %21s Len : %3d, Cap : %3d, ptr : %3p\n", " ", sb.Len(), sb.Cap(), &sb)
	return sb.String()
}

func base() {
	s := "this is a string"
	fmt.Printf("%v, %T\n", s, s) // this is a string, string

	fmt.Printf("%v, %T\n", s[2], s[2]) // 105, uint8. Because it should do the type convert first

	fmt.Printf("%v, %T\n", string(s[2]), string(s[2])) // i, string

	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b) // [116 104 105 115 32 105 115 32 97 32 115 116 114 105 110 103], []uint8
}
