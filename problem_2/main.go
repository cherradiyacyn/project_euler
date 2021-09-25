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

	limit, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if limit > max || limit < 2 {
		fmt.Printf("Retry with : 1 < n < %d\n", max)
		os.Exit(1)
	}

	sumOfEvens := func(limit int) int {
		var f1, f2, fn, sum int
		f2 += 1
		fn = f1 + f2
		for i := 0; fn < limit; i++ {
			if fn%2 == 1 {
				sum += fn
			}
			fn = f1 + f2
			f1 = f2
			f2 = fn
		}
		return sum
	}(limit)

	fmt.Printf("The sum of the even-valued terms in the Fibonacci sequence whose values do not exceed %d is : %d\n", limit, sumOfEvens)
}
