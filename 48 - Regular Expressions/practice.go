package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, err := regexp.Compile("p([a-z]+)ch") // 无 $ 符，将会是 true
	if err != nil {
		panic(err)
	}

	m1 := r.MatchString("peach sufix")
	fmt.Println("m1: ", m1)

	fmt.Println("r.Match([]byte(\"peach\")): ", r.Match([]byte("peach")))

	fmt.Println(r.FindString("peach punch"))
}
