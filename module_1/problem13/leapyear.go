package main

import "fmt"

func main() {
	var year int
	fmt.Scanln(&year)
	if year%4 == 0 || (year%400 == 0 && year%100 != 0) {
		fmt.Println(year, "is a leap year")
	} else {
		fmt.Println (year, "is not a leap year")
	}
}

// leap year or not
