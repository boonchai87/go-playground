package main

import "fmt"

type Person struct {
	name string
	age  int
}

// interface
type PersonInterace interface {
	Speak()
}

func main() {
	p := Person("dd", 5)
	fmt.Println(p)
}
