// 1
// 23
// 456
// 78910

package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	num := 1

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(num)
			num = num + 1
		}
		fmt.Println()
	}
}
