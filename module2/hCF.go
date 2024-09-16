// finding the hcf(highest common factore) also known as GCD
// largest positive integer that can divide both the numbers

package main

import (
	"fmt"
)

func main() {
	var n1, n2, n3 int
	fmt.Scanln(&n1, &n2, &n3)
	ans := 0
	ans1 := 0
	for i := 1; i <= n1 && i <= n2; i++ {
		if n1%i == 0 && n2%i == 0 {
			ans = i
		}
	}
	for i := 1; i <= ans && i <= n3; i++ {
		if ans%i == 0 && n3%i == 0 {
			ans1 = i
		}
	}
	fmt.Println(ans1)
}
