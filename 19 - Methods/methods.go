package main

import "fmt"

// Go supports methods defined on struct types
type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

// perimeter 周长
func (r * rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area())
	fmt.Println("pertimeter: ", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())
}
