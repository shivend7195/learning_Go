// conversion from to binary(0 and 1) to decimal(0-9)

package main

import (
	"fmt"
	"math"
)

func todeci(binary int64) int64 {
	var dec int64
	i := 0
	for binary != 0 {
		dig := binary % 10
		dec += dig * int64(math.Pow(float64(2), float64(i)))
		binary /= 10
		i++
	}
	return dec

}

func main() {
	var binary int64
	fmt.Scanln(&binary)
	ans := todeci(binary)
	fmt.Println(ans)

}
