package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	p := fmt.Println

	p(rand.Intn(100), ", ", rand.Intn(100))
	p("Float32: ", rand.Float32(), ", ", rand.Float32())
	p("Float64: ", rand.Float64(), ", ", rand.Float64())

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	p(r1.Intn(100), ", ", r1.Intn(100))

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	p(r2.Intn(100), ", ", r2.Intn(100))

	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	p(r3.Intn(100), ", ", r3.Intn(100))
}
