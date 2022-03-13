package main

import (
	"fmt"
	"math"
)

// Each bucket is used to paint 10 square meters and a bucket include 10 liter oil.
func paintBucket(width float64, height float64) (float64, error) {
	if width < 0 { // If width is invalid, return 0 and an error.
		return 0, fmt.Errorf("A width of %0.2f is invalid", width)
	}

	if height < 0 { // If height is invalid, return 0 and an error.
		return 0, fmt.Errorf("A width of %0.2f is invalid", height)
	}

	area := width * height

	// During painting, it should left more oil for repairing next time, so it should use Ceil(), not Floor().
	// Return the amount of paint, along with "nil", indicating there was no error.
	return math.Ceil(area / 10.0), nil
}

func bucketsNeeded(width float64, height float64) {
	// err is the second value to hold the second return value.
	amount, err := paintBucket(width, height)
	if err != nil {
		// Prints the error. If nil, there was no error.
		fmt.Println(err)
	} else {
		fmt.Printf("%0.2f buckets needed\n", amount)
	}
}

func main() {
	bucketsNeeded(4.2, -3.0)
	bucketsNeeded(5.2, 4.0)
}
