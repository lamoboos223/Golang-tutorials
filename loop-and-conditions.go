package main

import (
	"fmt"
)

func main() {

	limit := 3
	flag := false

	for i := 0; i < limit; i++ {
		if flag == true {
			fmt.Println(i)
		}
		flag = true
	}

	switch result := flag; result {
	case false:
		fmt.Println("it is false")
	case true:
		fmt.Println("it is true")
	default:
		fmt.Println("Not known state")
	}
}
