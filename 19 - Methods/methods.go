package main

import "fmt"

// Go supports methods defined on struct types
type rect struct {
	width, height int
}

// Methods can be fefined for either pointer or value receiver
// typeds.

// Here's an example of a pointer reviever
func (r *rect) area() int {
	return r.width * r.height
}

// perimeter 周长
// Here's an example of a value reviever
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func (r *rect) setWith(nWidth int) {
	r.width = nWidth
}

func (r rect) setHeight(nHeight int) {
	r.height = nHeight
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area())
	fmt.Println("pertimeter: ", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())

	// Go automatically handles conversion between values
	// and pointers for method calls. You may want to use a
	// pointer receiver type to avoid copying on method calls
	// or to allow the method to mutate the receiving struct.
	// 以下，只有结构的 width 会被改变
	r.setWith(1)
	r.setHeight(2)
	fmt.Println(r)

	rp.setWith(100)
	rp.setHeight(1000)
	fmt.Println(rp)
}
