package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const max = 37

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
	if input > max || input < 1 {
		fmt.Printf("Retry with : 1 < n < %d\n", max)
		os.Exit(1)
	}

	primeDecompositions := func(limit int) []map[int]int {
		sm := make([]map[int]int, 0)
		for i := 2; i <= limit; i++ {
			onePrimeDecomposition := func(number int) map[int]int {
				m := make(map[int]int)
				var isPrime bool
				for number > 1 {
					if isPrime {
						m[number]++
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
							m[pn]++
							number /= pn
							break
						}
						if pn > int(math.Sqrt(max)) {
							isPrime = true
						}
					}
					primesFile.Close()
				}
				return m
			}(i)
			sm = append(sm, onePrimeDecomposition)
		}
		return sm
	}(input)

	lcm := func(limit int, sm []map[int]int) int {
		m := make(map[int]int, 0)
		for i := 2; i <= limit; i++ {
			c := 0
			for _, m := range sm {
				if m[i] > c {
					c = m[i]
				}
			}
			if c != 0 {
				m[i] = c
			}
		}
		lcm := 1
		for k, v := range m {
			lcm *= int(math.Pow(float64(k), float64(v)))
		}
		return lcm
	}(input, primeDecompositions)

	fmt.Printf("The smallest positive number that is evenly divisible by all of the numbers from 1 to %d is : %d\n", input, lcm)
}

// https://youtu.be/2bIK1KkQ1k0
