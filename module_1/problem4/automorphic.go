// automorphic.go(square of a number ends with the same digit of number)

package main

import "fmt"

func automorphic(n int) int {
	sq := n * n
	for n != 0 {
		if n%10 != sq%10 {
			return 0
		}

		n /= 10
		sq /= 10
	}
	return 1
}

func main() {
	var n int
	fmt.Scanln(&n)
	if automorphic(n) == 1 {
		fmt.Println("IS AUTOMORPHIC")
	} else {
		fmt.Println("NOT AN AUTOMORPHIC")
	}
}
