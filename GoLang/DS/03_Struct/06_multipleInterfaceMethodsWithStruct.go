// Go program to illustrate the
// concept of multiple interfaces
package main

import "fmt"

// Interface
type Author interface {
	details()
	articles()
}

// Structure
type author struct {
	a_name string
	branch string
	college string
	year	 int
	salary int
	particles int
	tarticles int
}

// Implementing details() method of the interface
func (a author) details() {

	fmt.Printf("Author Name: %s", a.a_name)
	fmt.Printf("\nBranch: %s and passing year: %d", a.branch, a.year)
	fmt.Printf("\nCollege Name: %s", a.college)
	fmt.Printf("\nSalary: %d", a.salary)
	fmt.Printf("\nPublished articles: %d", a.particles)

}

// Implementing articles() method of the interface
func (a author) articles() {

	pendingarticles := a.tarticles - a.particles
	fmt.Printf("\nPending articles: %d", pendingarticles)
}

func main() {

	// Assigning values to the structure
	values := author{
		a_name: "Mickey",
		branch: "Computer science",
		college: "XYZ",
		year:	 2012,
		salary: 50000,
		particles: 209,
		tarticles: 309,
	}

	// Accessing the methods of the interface
	var i Author = values
	i.details()
	i.articles()
}
