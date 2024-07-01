package main

import (
	"fmt"
)

func main() {
	array1 := []int{1, 2, 3}
	var array2 [3]int
	array2[0] = 1
	array2[1] = 2
	array2[2] = 3

	array3 := [3]int{1, 2, 3}

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)
}
