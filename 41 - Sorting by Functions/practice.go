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

// 声明一个复杂的结构体
type combination struct {
	order   int
	other   string
	complex privateStr
}

type cslice []combination

func (s cslice) Len() int {
	return len(s)
}

func (s cslice) Less(i, j int) bool {
	return s[i].order < s[j].order
}

func (s cslice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(privateStr(fruits))
	fmt.Println(fruits)

	unSort := []combination{
		{
			order:   33,
			other:   "hello",
			complex: []string{"peach", "banana", "kiwi"},
		},
		{
			order: 21,
			other: "world",
		},
		{
			order: 48,
			other: "bill",
		},
	}
	sort.Sort(cslice(unSort))
	fmt.Println(unSort)
}
