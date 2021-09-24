package main

import (
	"fmt"
	"os"
	"strconv"
)

func reverse(s string) string {
	bs := make([]byte, 0)
	for i := len(s) - 1; i >= 0; i-- {
		bs = append(bs, s[i])
	}
	return string(bs)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "No arguments...")
		os.Exit(1)
	}

	ndigits, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if ndigits > 3 || ndigits < 2 {
		fmt.Printf("Retry with : 1 < n < 4\n")
		os.Exit(1)
	}

	number := func(len int) int {
		bs := make([]byte, 0)
		for i := len; i > 0; i-- {
			bs = append(bs, '9')
		}
		n, _ := strconv.Atoi((string(bs)))
		return n
	}(ndigits)

	largestPalindrome := func(limit int) string {
		var lp int
		for i := 0; i <= limit; i++ {
			for j := 0; j <= limit; j++ {
				product := i * j
				w := strconv.Itoa(product)
				if w == reverse(w) && product > lp {
					lp = product
				}
			}
		}
		return strconv.Itoa(lp)
	}(number)

	fmt.Printf("The largest palindrome made from the product of two %d-digit numbers is : %s\n", ndigits, largestPalindrome)
}
