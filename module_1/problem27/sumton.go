package main

import "fmt"

func main() {

	var num int
	fmt.Scanln(&num)
	sum := 0
	for num > 0 {
		ls := num % 10
		sum += ls
		num /= 10
	}
	fmt.Println(sum)
}

// sum upto n number
