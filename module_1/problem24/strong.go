// if a number is strong(a number in which the sum of the factorial of individual digits of the numbers is equal to the number itself.)
package main

import "fmt"

func factorials(n int) int {
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	return fact
}

func main() {
	var n int
	sum := 0
	fmt.Scanln(&n)
	new := n
	for n > 0 {
		ls := n % 10
		sum += factorials(ls)
		n /= 10
	}
	if sum == new {
		fmt.Println("STRONG NUMBER")
	} else {
		fmt.Println("NOT STRONG NUMBER")
	}

}
