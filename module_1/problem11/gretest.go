package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanln(&a, &b, &c)
	if a > b && a > c {
		fmt.Println(a, "is greatest")
	} else if b > c && b > a {
		fmt.Println(b, "is greatest")
	} else {
		fmt.Println(c, "is greatest")
	}
}

//gretest of three
