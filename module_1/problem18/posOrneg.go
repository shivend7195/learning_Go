package main

import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)
	if num >= 0 {
		fmt.Println("POSITIVE")
	} else {
		fmt.Println("NEGATIVE")
	}
}

// positive or negative
