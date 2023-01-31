package main

import (
	"fmt"
	"sort"
)

type privateStr []string

func (s privateStr) Len() int {
	return len(s)
}

func (s privateStr) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func (s privateStr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(privateStr(fruits))
	fmt.Println(fruits)
}
