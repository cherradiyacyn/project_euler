package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

const max = math.MaxInt16

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
	if limit > max || limit < 1 {
		fmt.Printf("Retry with : 0 < n < %d\n", max)
		os.Exit(1)
	}

	primes := []int{2}
	next := primes[len(primes)-1] + 1

	for c := len(primes); c < limit; next += 2 {

		isPrime := func(n int, primes []int) bool {
			for _, p := range primes {
				if n%p == 0 {
					return false
				}
			}
			return true
		}(next, primes)

		if isPrime {
			primes = append(primes, next)
			c++
		}
	}

	fmt.Printf("The %dth prime number is : %d\n", limit, primes[len(primes)-1])
}
