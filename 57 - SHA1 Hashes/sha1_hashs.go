package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 this string"
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	fmt.Printf("\"%s\" sha1 get: %x\n", s, bs)
}
