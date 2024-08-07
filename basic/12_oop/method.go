package main

import (
	"fmt"
	"math"
)

type Object struct {
	X, Y float64
}

func (v Object) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func Abs(v Object) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Object{3, 4}
	fmt.Println(v.Abs())
}
