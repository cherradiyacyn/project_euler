package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// the largest number in "prime_numbers.txt" is 989999; 989999² = 980098020001
// max <= 980098020001
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
	if input > max || input < 2 {
		fmt.Printf("Retry with : 1 < n < %d\n", max)
		os.Exit(1)
	}

	largestPrimeFactor := func(number int) int {
		var lpf int
		var isPrime bool
		for number > 1 {
			if isPrime {
				lpf = number
				break
			}
			primesFile, err := os.Open("prime_numbers.txt")
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
			scanr := bufio.NewScanner(primesFile)
			for scanr.Scan() {
				pn, _ := strconv.Atoi(scanr.Text())
				if number%pn == 0 {
					lpf = pn
					number /= pn
					break
				}
				if pn > int(math.Sqrt(max)) {
					isPrime = true
				}
			}
			primesFile.Close()
		}
		return lpf
	}(input)

	fmt.Printf("The largest prime factor of the number %d is : %d\n", input, largestPrimeFactor)
}
