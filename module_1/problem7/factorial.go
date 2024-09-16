// factorial number

package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	fact := 1
	if n < 0 {
		fmt.Println("NOT POSSIBLE")
	} else {
		for i := 1; i <= n; i++ {
			fact *= i
		}
		fmt.Println(fact)
	}
}

// OR 

// package main

// import "fmt"

// func factorial(n int) int {
// 	if n == 0 {
// 		return 1
// 	}

// 	return n * factorial(n-1)
// }
// func main() {
// 	var n int
// 	fmt.Scanln(&n)

// 	fact := factorial(n)
// 	fmt.Println(fact)
// }
