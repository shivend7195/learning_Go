//    A
//   ABA
//  ABCBA
// ABCDCBA

package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		char := 'A'
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}

		breakPoint := (2*i + 1) / 2

		for j := 1; j <= (2*i + 1); j++ {
			fmt.Print(string(char))
			if j <= breakPoint {
				char++
			} else {
				char--
			}
		}

		fmt.Println()
	}
}
