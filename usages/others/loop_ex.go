package main

import "fmt"

func main() {
	forTypeOne()
	forTypeTwo()
	forAndIfUsage()
	forSwitchUsage()
}

func forTypeOne() {
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}
}

func forTypeTwo() {
	for i := 5; i <= 10; i++ {
		fmt.Println(i)
	}
}

func forAndIfUsage() {
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}
}

func forSwitchUsage() {
	for i := 1; i <= 6; i++ {
		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")
		default:
			fmt.Println("Unknown Number")
		}
	}
}
