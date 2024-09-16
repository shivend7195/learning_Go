//	  *
//	 ***
//	*****

// printing space and star but not  printing space after star
package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ { // look for relation in space
			fmt.Print(" ")
		}
		for k := 0; k < (2*i)+1; k++ { // for stars
			fmt.Print("*")
		}

		fmt.Println()

	}
}
