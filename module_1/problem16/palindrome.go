// palindrome ( if the reverse of the number and the numbers itself are equal )
package main

import "fmt"

func main() {
	var num int
	rev := 0
	fmt.Scanln(&num)
	new := num
	for num > 0 {
		ls := num % 10
		rev = rev*10 + ls
		num /= 10
	}
	if rev == new {
		fmt.Println("PALINDROME")
	} else {
		fmt.Println("NOT PALINDROME")
	}

}
