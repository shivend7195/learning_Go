// abundunt number (is the sum of factors less than the number itself is more than the number itself)

package main

import "fmt"

func abund(n int) int {
	sum := 0
	for i := 1; i*i < n; i++ {
		if i == 1 {
			sum += i
		} else if i == n/i {
			sum += i
		} else {
			sum += (sum + i + n/i)
		}
	}
	return sum
}

func main() {
	var n int
	fmt.Scanln(&n)
	if abund(n) > n {
		fmt.Println("ABUNDUNT NUMBER")
	} else {
		fmt.Println("NOT AN ABUNDUNT NUMBER")
	}

}
