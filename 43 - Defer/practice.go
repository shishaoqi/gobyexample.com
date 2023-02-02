package main

import (
	"fmt"
	"os"
)

// 不足：
//
func main() {
	fname := "./defer.txt"
	f := creatFile(fname)
	defer closeFile(f)
	writeFile(f)
}

func creatFile(f string) *os.File {
	fmt.Println("create file")
	fl, err := os.Create(f)
	if err != nil {
		fmt.Println(err)
	}
	return fl
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "practice data")
}

func closeFile(f *os.File) {
	fmt.Println("close file")
	err := f.Close()
	if err != nil {
		os.Exit(1)
	}
}
