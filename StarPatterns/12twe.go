// 1
// 0 1
// 1 0 1
// 0 1 0 1

package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if i == j || (i%2 == 0 && j%2 == 0) || (i%2 != 0 && j%2 != 0) {
				fmt.Print("1")
			} else {
				fmt.Print("0")
			}
		}
		fmt.Println()
	}
}
