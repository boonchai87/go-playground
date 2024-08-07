package main

import (
	"fmt"
)

func main() {
	//var j int = 0
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	/* infinite loop
	for {
		fmt.Println("fuck")
	}*/
}
