// ocatal to decimal

package main

import (
	"fmt"
	"math"
)

func to_deci(ocatal int64) int64 {
	var dec int64
	i := 0
	for ocatal != 0 {
		dig := ocatal % 10
		dec += dig * int64(math.Pow(float64(8), float64(i)))
		ocatal /= 10
		i++
	}
	return dec

}

func main() {
	var ocatal int64
	fmt.Scanln(&ocatal)
	ans := to_deci(ocatal)
	fmt.Println(ans)

}
