package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Scanln(&a, &b)
	ans := ((b*(b+1))/2 - (a*(a+1))/2) + a
	fmt.Println(ans)

}

// sum within range 2 and 4
// 2+3+4 = 9
