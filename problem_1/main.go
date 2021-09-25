package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

const max = math.MaxInt32

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "No arguments...")
		os.Exit(1)
	}

	input, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if input > max || input < 3 {
		fmt.Printf("Retry with : 2 < n < %d\n", max)
		os.Exit(1)
	}

	sumOfMultiples := func(limit int) int {
		s := 0
		for i := 1; i < limit; i++ {
			if i%3 == 0 || i%5 == 0 {
				s += i
			}
		}
		return s
	}(input)

	fmt.Printf("sum of all the multiples of 3 or 5 below %d is : %d\n", input, sumOfMultiples)
}
