// prime factors of a number

package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	for n%2 == 0 {
		fmt.Print(2, " ")
		n /= 2
	}
	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			fmt.Print(i, " ")
			n /= i
		}
	}
	if n > 2 {
		fmt.Print(n)
	}

}
