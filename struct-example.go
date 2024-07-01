package main

import (
	"fmt"
)

type Person struct {
	age  int
	name string
}

func main() {
	x := Person{30, "Ahmed"}
	var y = Person{24, "Sara"}
	z := &x
	z.age = 29
	w := x
	w.age = 28

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(*z)
	fmt.Println(w)
}
