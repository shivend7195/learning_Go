package main

import (
	"fmt"
	"math"
)

func main() {
	var base, exponent float64
	fmt.Scanln(&base, &exponent)
	ans := float64(math.Pow(float64(base), float64(exponent)))
	fmt.Println(ans)
}
