// Program to use function as a field  of struct

package main
import "fmt"

// initialize the function Rectangle
type Rectangle func(int, int) int

// create structure
type Shapes struct {
  length  int
  breadth int

  // function as a field of struct
  rect Rectangle
}

func main() {

  // assign values to struct variables
  s := Shapes{
    length:  10,
    breadth: 20,
    rect: func(length int, breadth int) int {
      return length * breadth
    },
  }

  fmt.Println("Area of Rectangle: ", s.rect(s.length, s.breadth))
}