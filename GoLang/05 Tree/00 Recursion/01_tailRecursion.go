package main

import "fmt"


func tailRecursion(num int) {
	if num <= 0 {
		return
	}
	
	fmt.Print("->", num)

	tailRecursion(num - 1)
}

func main() {
	tailRecursion(10)
}