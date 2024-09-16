package main

import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)
	if num & 2 == 0 {
		fmt.Println("EVEN")
	} else {
		fmt.Println("ODD")
	}
}

// even or odd
