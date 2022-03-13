package main

import (
	"fmt"
)

func main() {
	/*
		%s and %d are verbs, they will be substituted with a value in a particular format.
		In the below example, "gumballs" will replace %s, 23 will replace $d.

		Rules are following:
			%f 	means Floating-point number
			%d 	means Decimal integer
			%s  means string
			%t  means Boolean
			%v  means Any value (Won't format anything. for example "\t" will be "	".)

			%#v means Any value, formatted as it would appear in Go code. reveals an empty string, a tab character, and a newline, all of which were invisible when printed with %v.

			%T  means Type of the supplied value.
			%%  means A lateral percent sign. It will output '%'
	*/
	fmt.Printf("The %s cost %d cents each.\n", "gumballs", 23) // The gumballs cost 23 cents each.

	fmt.Printf("A float: %f\n", 3.141)
	fmt.Printf("An integer: %d\n", 24)
	fmt.Printf("A String: %s\n", "hello world")
	fmt.Printf("A boolean: %t\n", true)
	fmt.Printf("Values: %v %v %v %v\n", 1.5, "\t", "\n", false)
	fmt.Printf("Values: %#v %#v %#v %#v\n", 1.8, "\t", "\n", true) // Values: 1.8 "\t" "\n" true
	fmt.Printf("Types: %T %T %T\n", 2.3, "\t", true)               // Types: float64 string bool
	fmt.Printf("Percent sign: %%\n")

	// Using %f in go to format the amount of paint needed.
	// Formatting verbs let you specify the width of the formatted value.

	// fmt.Printf("%4s | %s\n", "Product", "Cost in Cents")	//It will see the diff between %12s and %4s
	fmt.Printf("%12s | %s\n", "Product", "Cost in Cents")
	fmt.Println("-----------------------------")
	fmt.Printf("%12s | %2d\n", "Stamps", 50)
	fmt.Printf("%12s | %2d\n", "Paper Clips", 5)
	fmt.Printf("%12s | %2d\n", "Tape", 99)

	/*
		Output is below:

			 Product | Cost in Cents
		-----------------------------
		      Stamps | 50
		 Paper Clips |  5
		        Tape | 99
	*/

	/*
		Formatting fractional number widths.
		%x.yf, here x and y mean different numbers.

		The % is start of the formatting verb.
		The x is the minimum width of entire number. The minimum width of the entire number includes decimal places and the decimal point.
		The y is the width after decimal point
		The f means formatting verb type
	*/

	fmt.Printf("%%7.3f: %7.3f\n", 12.3456) // 12.346
	fmt.Printf("%%7.3f: %7.3f\n", 12.3453) // 12.345
	fmt.Printf("%%7.3f: %7.2f\n", 12.3456) // 12.35
	fmt.Printf("%%7.3f: %7.1f\n", 12.3456) // 12.3
	fmt.Printf("%%.1f: %.1f\n", 12.3456)   // 12.3, ".f means left all int part."
	fmt.Printf("%%.2f: %.2f\n", 12.3456)   // 12.35, ".f means left all int part."
}
