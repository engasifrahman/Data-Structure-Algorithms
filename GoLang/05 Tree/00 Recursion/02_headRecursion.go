package main

import "fmt"


func tailRecursion(num int) {
	if num <= 0 {
		return
	}

	tailRecursion(num - 1)

	fmt.Print("->", num)
}

func main() {
	tailRecursion(10)
}