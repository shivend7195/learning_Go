// *
// **
// ***
// ****
// ***
// **
// *

package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 0; i < n; i++ {
		for j := n - 1; j > i; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
