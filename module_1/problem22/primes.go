package main

import (
	"fmt"
)

func primes(n int) bool {
	if n < 2 {
		return false
	} else {
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				return false
			}
		}
	}
	return true
}

func main() {
	var low, high int
	fmt.Scan(&low, &high)
	for i := low; i <= high; i++ {
		if primes(i) {
			fmt.Print(i, " ")
		}
	}
}

// primes within range
