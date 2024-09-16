package main

import (
	"fmt"
)

func primeFact(n int) []int {
	fact := make([]int, 0)
	for n%2 == 0 {
		fact = append(fact, 2)
		n /= 2
	}
	for i := 3; i*i <= n; i += 2 {
		for n%i == 0 {
			fact = append(fact, i)
			n /= i
		}
	}
	if n > 2 {
		fact = append(fact, n)
	}
	return fact
}

func main() {
	var n int
	fmt.Scan(&n)
	factors := primeFact(n)
	fmt.Println(factors)
}

// binary factors.
