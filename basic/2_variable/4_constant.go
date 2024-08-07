package main

import (
	"fmt"
)

const PI = 3.14

// Constants cannot be declared using the := syntax.
const Truth bool = true

func main() {
	fmt.Println(PI)
	fmt.Println(Truth)
}
