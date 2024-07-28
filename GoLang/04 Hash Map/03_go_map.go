package main
import "fmt"

func main() {
	// create a map
	personAge := map[string]int{"Hermione": 21, "Harry": 20, "John": 25}
	fmt.Println("Initial Map: ", personAge)

	// add element with the key Robin
	personAge["Robin"] = 18

	// remove element of map with key John
	delete(personAge, "John")

	fmt.Println("Updated Map: ", personAge)

	// for-range loop to iterate through each key-value of map
	for name, age := range personAge {
		fmt.Printf("%s is %d years old\n", name, age)
	}
}