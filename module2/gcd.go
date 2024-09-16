// gcd (greatest common divisor) of two number

package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Scanln(&n1, &n2)

	for n1 != n2 {
		if n1 > n2 {
			n1 -= n2
		} else {
			n2 -= n1
		}
	}
	fmt.Println(n1)
}

// above method  Eculidean Alogrithm : GCD of two numbers remains constant even if we substract the smaller nubmer
// from the higher number
