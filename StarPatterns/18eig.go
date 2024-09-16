// D
// C D
// B C D
// A B C D

package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		for ch := 'A' + rune(n-1) - rune(i); ch <= 'A'+rune(n-1); ch++ {
			fmt.Print(string(ch) + " ")
		}
		fmt.Println()
	}
}
