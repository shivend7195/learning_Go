package main

import (
	"fmt"
	"math"
)

func main() {
	var num int
	fmt.Scanln(&num)
	if num <= 1 {
		fmt.Println("NA")
	} else {
		limit := int(math.Sqrt(float64(num)))
		//limit := int(math.Sqrt(float64(num)))
		fl := 0
		for i := 2; i <= limit; i++ {
			if num%i == 0 {
				fl = 1
			}
		}
		if fl == 0 {
			fmt.Println("PRIME")
		} else {
			fmt.Println("NOT PRIME")
		}
	}
}

// no is prime or not
