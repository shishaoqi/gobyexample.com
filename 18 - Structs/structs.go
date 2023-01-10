package main

import "fmt"

type person struct {
	name string
	age  int
}

// You can safely return a pointer to local variable
// as a local variable will servive the cope of the function
func newPerson(name string) *person {
	p := person{
		name: name,
	}
	p.age = 30
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})              // This syntax creates a new struct
	fmt.Println(person{name: "Alice", age: 30}) // You can name the fields when initializing a struct
	fmt.Println(person{name: "Fred"})           // Omitted fields will be zero-valued
	fmt.Println(&person{name: "Andy", age: 40}) // An & prefix yields a pointer to the struct

	// It’s idiomatic to encapsulate new struct creation in constructor functions
	fmt.Println(newPerson("jon"))
	ps := person{name: "Sean", age: 50}
	fmt.Println(ps.name) // Access struct fields with a dot
	psPtr := &ps
	fmt.Println(psPtr.age) // You can also use dots with struct pointer - the pointers are automatically derefenced

	psPtr.age = 51
	fmt.Println(psPtr.age) // structs are mutable

}
