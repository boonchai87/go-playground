package main

import "fmt"

type Vertex3 struct {
	X int
	Y int
}

var (
	// constructor
	v1 = Vertex3{1, 2}  // has type Vertex
	v2 = Vertex3{X: 1}  // Y:0 is implicit
	v3 = Vertex3{}      // X:0 and Y:0
	p  = &Vertex3{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(Vertex3{1, 2})
	fmt.Println(v1, p, v2, v3)
}
