// two num are freindly pair (if (sum of divisors)/num1  == (sum of divisor)/num2)
package main

import "fmt"

func sumOfDig(n int) int {
	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	var num1, num2 int
	fmt.Scanln(&num1, &num2)
	sum1 := sumOfDig(num1)
	sum2 := sumOfDig(num2)
	ratio1 := sum1 / num1
	ratio2 := sum2 / num2

	if ratio1 == ratio2 {
		fmt.Println("FRIENDLY PAIR")
	} else {
		fmt.Println("NOT FRIENDLY PAIR")
	}
}
