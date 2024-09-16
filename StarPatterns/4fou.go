// 1
// 1 2
// 1 2 3
// 1 2 3 4
package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}
