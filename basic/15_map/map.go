package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex = make(map[string]Vertex)

var m2 = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	//m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	fmt.Println(m2)

	m3 := make(map[string]int)

	m3["Answer"] = 42
	fmt.Println("The value:", m3["Answer"])

	m3["Answer"] = 48
	fmt.Println("The value:", m3["Answer"])

	delete(m3, "Answer")
	fmt.Println("The value:", m3["Answer"])

	value, ok := m["Answer"]
	fmt.Println("The value:", value, "Present?", ok)

	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	fmt.Printf("a\t%v\n", a)
}
