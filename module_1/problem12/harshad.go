// harshad number(a number is divisible by sum of individual digit of a number)

package main

import "fmt"

func sumOfDigit(n int) int {
	sum := 0
	for n > 0 {
		ls := n % 10
		sum += ls
		n /= 10
	}
	return sum
}

func main() {
	var n int
	fmt.Scanln(&n)
	if n%sumOfDigit(n) == 0 {
		fmt.Println("HARSHAD")
	} else {
		fmt.Println("NOT HARSHAD")
	}
}
