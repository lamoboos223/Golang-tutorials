package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"test1": {9.8, 3.5},
}

func main() {
	fmt.Println(m)
	a := Vertex{1.2, 5.1}
	c := map[string]Vertex{
		"test2": a,
	}
	fmt.Println(c)

	m["test1"] = a
	fmt.Println(m)
}
