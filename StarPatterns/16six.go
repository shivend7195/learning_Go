// A
// BB
// CCC
// DDDD

package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	for i := 0; i <= n; i++ {
		char := rune('A' + i)
		for j := 0; j <= i; j++ {
			fmt.Print(string(char))

		}
		fmt.Println()

	}
}
