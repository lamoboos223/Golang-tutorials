package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName, lastName string
}
type Vertex struct {
	X, Y float64
}

func add(x, y int) (string, int) {
	return "The result is:", x + y
}

func add2(x, y int) (msg string, result int) {
	msg = "The result is:"
	result = x + y
	return
}

func (p Person) printFullName() string {
	return p.firstName + " " + p.lastName
}
func (p *Person) capitalize() {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}
func main() {
	fmt.Println(add(1, 1))
	fmt.Println(add2(1, 1))

	person := Person{"lama", "alosaimi"}

	fmt.Println(person.printFullName())
	person.capitalize()
	fmt.Println(person)

}
