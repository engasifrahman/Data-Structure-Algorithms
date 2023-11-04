package main

import "fmt"

// Recursive function
func fun(n int) {
	// fmt.Printf("%d ", n)

    if n > 0 {
		fmt.Printf("%d ", n)

        // Calling once
        fun(n - 1)

        // Calling twice
        fun(n - 1)
    }
}

func main() {
    fun(3)
}
