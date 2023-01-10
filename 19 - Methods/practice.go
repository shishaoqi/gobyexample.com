package main

import "fmt"

type rect struct {
	width  float64
	height float64
}

func (r *rect) area() float64 {
	return r.width * r.height
}

func (r *rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func newRect(w, h float64) *rect {
	r := rect{
		width:  w,
		height: h,
	}

	return &r
}

func (r *rect) setHeight(h float32) {
	th := float64(h)
	r.height = th
}

func main() {
	var rt *rect
	rt = newRect(1.0, 2.0)
	fmt.Printf("The rectangle's area is: %v and perimeter is: %v\n", rt.area(), rt.perimeter())

	fmt.Println(rt)
	rt.setHeight(4.5)
	fmt.Printf("Updating rectangle's height, rt = %v\n", rt)
	fmt.Printf("After updated, the rectangle's area is: %v and perimeter is: %v\n", rt.area(), rt.perimeter())
}
