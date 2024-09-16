// reverse of a number

package main

import "fmt"

func main() {
	var num int
	rev := 0
	fmt.Scanln(&num)
	for num > 0 {
		ls := num % 10
		rev = rev*10 + ls
		num /= 10
	}
	fmt.Println(rev)

}
