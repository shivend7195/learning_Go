// ABCD
// ABC
// AB
// A

package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		for j := 'A'; j <= rune('A'+(n-i-1)); j++ {
			fmt.Print(string(j))
		}
		fmt.Println()
	}
}
