// ****
// ****
// ****
// ****

package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ { //as rows and coulmns are equal so as the row proceeds the columns do too
			fmt.Print("*")
		}
		fmt.Println()
	}

}
