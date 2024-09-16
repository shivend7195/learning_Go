// lcm of three numbers (lowest common multiple)

package main

import (
	"fmt"
	"math"
)

func main() {
	var n1, n2, n3 int
	fmt.Scanln(&n1, &n2, &n3)
	ans := 0
	ans1 := 0
	max := int(math.Max(float64(n1), float64(n2)))
	for i := max; i <= (n1 * n2); i++ {
		if i%n1 == 0 && i%n2 == 0 {
			ans = i
			break
		}
	}
	max1 := int(math.Max(float64(ans), float64(n3)))
	for i := max1; i <= (ans * n3); i++ {
		if i%ans == 0 && i%n3 == 0 {
			ans1 = i
			break
		}
	}
	fmt.Println(ans1)
}
