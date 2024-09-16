// sum upto n
package main

import "fmt"

func main() {
	var num int
	sum := 0
	fmt.Scanln(&num)
	sum += (num * (num + 1)) / 2
	fmt.Println(sum)
}
