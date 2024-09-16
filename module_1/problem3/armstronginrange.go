// armstrong in range

package main

import (
	"fmt"
	"math"
)

func isArmstrong(num int) bool {
	dig := 0
	temp := num
	for temp > 0 {
		dig++
		temp /= 10
	}
	temp = num
	sum := 0
	for temp > 0 {
		ls := temp % 10
		sum += int(math.Pow(float64(ls), float64(dig)))
		temp /= 10
	}
	return sum == num
}

func main() {
	var low, high int
	fmt.Scanln(&low, &high)

	for i := low; i <= high; i++ {
		if isArmstrong(i) {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
