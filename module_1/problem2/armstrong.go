// abcd == a^4+b^4+c^4+d^4

package main

import (
	"fmt"
	"math"
)

func main() {
	var num int
	var len int
	fmt.Scanln(&num)
	new := num
	number := num
	for num > 0 {
		len++
		num /= 10
	}
	var ls int
	var sum int
	for new > 0 {
		ls = new % 10
		sum += int(math.Pow(float64(ls), float64(len)))
		new /= 10
	}
	if sum == number {
		fmt.Println("ARMSTRONG")
	} else {
		fmt.Println("NOT ARMSTRONG")
	}

}
