package main

import (
	"fmt"
	"sort"
	"strings"
)

type Stringer interface {
	Display() string // Any type is fmt.Stringer, if it has a String() method that returns a string.
}

func showCofferPot(s Stringer) string {
	return s.Display()
}

type CoffeePot string

func (c CoffeePot) Display() string { // Satisfy the Stringer interface.
	return string(c) + " coffee pot" // Method needs to return a string
}

func sortString() {
	message := []string{"D", "E", "F", "X", "S", "Y", "Z", "G", "H", "I", "J", "K", "V", "W", "L", "M", "N", "O", "P", "B", "C", "Q", "R", "A", "T", "U"}
	sort.Strings(message)
	fmt.Printf("Sorted string is: %v \n", strings.Join(message, "")) // ABCDEFGHIJKLMNOPQRSTUVWXYZ
}

func printStringType() {
	// fmt.Print will output the space in strings.
	// So if the output is a string, it should have a space. If not, it will connect with other outputs.
	// If not a string, it will include a space automatically.
	fmt.Print("fmt.Print output:\n", 1, nil, false, 3, true, "\ntest:", "not have a space\n", "test:", " have a space\n")

	// fmt.Println will default output a space in any output type.
	// So if it is a string, don't need a space. If it includes a space, it will also output it.
	fmt.Println("fmt.Println output:\n", 1, nil, false, 3, true, "string")

	// Sprint does everything what Print does but does not write the resulting string to the standard output,
	// but it returns the resulting string.
	fmt.Print("fmt.Sprint and fmt.Sprintf output is nothing, but it can be used set value.\n")
	s1 := fmt.Sprintf("It is a %s ", "string\n")
	s2 := fmt.Sprintln("fmt.Println output:\n", 1, nil, false, 3, true, "string")
	fmt.Print(s1, s2)

	formatVerbs()
}

func formatVerbs() {
	test := struct {
		name string
		age  int
	}{"John", 26}

	// %v formatting verb is only output value.
	fmt.Printf("v is output %v\n", test)

	// %#v formatting verb is output value and key, if the output has a key.
	fmt.Printf("#v is output %#v\n", test)

	// Display field names in a struct or pointer to a struct, use %+v.
	fmt.Printf("%+v \n", test)
	fmt.Printf("%+v \n", &test)

	// %T formatting verb represents the data-type of a value
	fmt.Printf("%T \n", test)

	// %t formatting verb is used to format boolean value into a string word true or false.
	fmt.Printf("%t \n", false)

	// %d is base10 format.
	fmt.Printf("Dexadecimal 0xFF converts to base10 is %d\n", 0xFF) // hxadecimal or dexadecimal, all means 16进制
	fmt.Printf("octal 011 converts to base10 is %d\n", 011)         // 8进制
	fmt.Printf("rune # converts to base10 is %d\n", '#')            // 特殊符号
	fmt.Printf("output sign in base10 is %+d\n", 11)                // 运算符
	fmt.Printf("output sign in base10 is % d\n", 11)                // leave space when sign is not mentioned.

	// %b is base2/binary number format.
	fmt.Printf("9 in base2 is %b\n", 9)
	fmt.Printf("0xF in base2 is %b\n", 0xF)

	// %x (hxadecimal format) and $o (octal format)
	// %f (decimal format)
	// %e (scientific notation format)
	// %c (character format)
	// %U (Unicode format)

	// %s (string format)
	fmt.Printf("Usual string format is %s\n", "hello")
	// %q (escaped string format)
	fmt.Printf("Excaped string format is %q\n", "hello")
}

func main() {
	sortString()

	result := showCofferPot(CoffeePot("LuxBrew"))
	fmt.Println(result)

	printStringType()
}
