// *****  0 5
//  ***   1 3
//   *    2 1

package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*n-(2*i+1); j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
