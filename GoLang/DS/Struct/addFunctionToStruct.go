package main

import "fmt"

// taking a struct
type Rectangle struct {
	length, breadth int
}

func (rect Rectangle) Area() int {
	return rect.length * rect.breadth
}

func main() {
	rect := Rectangle{10, 12}
	fmt.Println("Length and Width are:", rect)
	fmt.Println("Area of Rectangle: ", rect.Area())
}
