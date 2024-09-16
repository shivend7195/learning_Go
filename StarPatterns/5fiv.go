// 1
// 22
// 333
// 4444
// rows are going as whole and coulumn is less than row
package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(i)
		}
		fmt.Println()
	}
}
