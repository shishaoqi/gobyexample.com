package main

import (
	"fmt"
	"sort"
)

func main() {
	var ints []int = []int{9, 4, 5}
	sort.Ints(ints)
	fmt.Println(ints)

	var strs []string = []string{"b", "d", "A", "w"}
	sort.Strings(strs)
	fmt.Println(strs)

	fmt.Println("The ints is sorted: ", sort.IntsAreSorted(ints))
}
