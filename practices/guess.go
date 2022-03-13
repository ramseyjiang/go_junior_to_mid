package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"

	// "reflect"
	"strconv"
	"strings"
	"time"
)

func main() {
	/*
		Generate a current date and time, which can be used to get a different seed value every time.

		The rand.Seed function expects an integer, so we can’t pass it a Time value directly.
		The time.Now() will get a time.time type value.
		So we call the Unix method on the Time, which will convert it to an integer.
	*/
	seconds := time.Now().Unix()

	/*
		To get different random numbers, we need to pass a value to the rand.Seed function.
		That will “seed” the random number generator—that is, give it a value that it will use to generate other random values.
		But if we keep giving it the same seed value, it will keep giving us the same random values,
	*/

	// Seed the random number generator.
	rand.Seed(seconds) // If not seed the random number, we always get the same value.

	/*
		Generate a random num from 0 to 100.
		rand.Intn(100) will generate a num in (0, 99), that's why here have "+1".
		Otherwise, here can use rand.Intn(101)
	*/
	randomNum, times := rand.Intn(100)+1, 10
	fmt.Println("Computer has generated a random number between 1 and 100. Can you guess it?")

	// Create a bufio.Reader which lets us read keyboard input
	reader := bufio.NewReader(os.Stdin)

	for times > 0 {
		fmt.Println("Please guess a number, you have ", times, "times left")
		fmt.Print("Make a guess:")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)

		// Convert the input string to an integer.
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(guess, reflect.TypeOf(guess), randomNum, guess > randomNum)
		times-- // '--', they can only be put in at the back of the variable, if put them in front, it will have errors.
		if guess > randomNum {
			fmt.Println("Your guess number was gather than the target, you have", times, "left.")
			continue
		} else if guess < randomNum {
			fmt.Println("Your guess number was less than the target, you have", times, "left.")
			if times == 0 && guess != randomNum {
				fmt.Println("Sorry, you used all times you could guess, the target is", randomNum, ",please play again.")
			}
			continue
		} else {
			fmt.Println("Your guess number equals to the target, you win! You used", 10-times, "times.")
			break
		}
	}
}
