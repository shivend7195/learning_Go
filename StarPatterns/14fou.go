// A
// AB
// ABC
// ABCD

package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		for char := 'A'; char <= rune('A'+i); char++ {
			fmt.Print(string(char))
		}
		fmt.Println()
	}
}
