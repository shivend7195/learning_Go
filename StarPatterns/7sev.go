// 1234
// 123
// 12
// 1

package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i+1; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}
