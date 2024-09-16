// a perfect numbe(a number in which the sum of the proper positive divisors of the number is equal to the number itself.)

package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	new := n
	sum := 0
	for i := 1; i < n; i++ {
		if new%i == 0 {
			sum += i
		}
	}
	if sum == n {
		fmt.Println("PERFECT NUMBER")
	} else {
		fmt.Println("NOT PERFECT NUMBER")
	}
}
