// nth term of fibonacchi term

package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	a := 0
	b := 1
	var nextTerm int
	for i := 0; i < n; i++ {
		nextTerm = a + b
		a = b
		b = nextTerm
	}
	fmt.Println(nextTerm)

}
