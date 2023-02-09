package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println("resource data: ", true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println("resource data: ", 1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println("resource data: ", 2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println("resource data: ", "gopher")
	fmt.Println(string(strB)) // marshal后，字符串带有双引号

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println("------\nresource array: ", slcD)
	fmt.Println(string(slcB)) // 每个元素字符串带有双引号
	slcF := []string{}
	json.Unmarshal(slcB, &slcF)
	fmt.Println("Unmarshal array: ", slcF)

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println("------\nresource map: ", mapD)
	fmt.Println(string(mapB))
	mapF := make(map[string]int)
	json.Unmarshal(mapB, &mapF)
	fmt.Println("Unmarshal map: ", mapF)

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println("------\nresource response1: ", res1D)
	fmt.Println(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println("------\nresource response2: ", res2D)
	fmt.Println(string(res2B))

	fmt.Println("---------- Unmarshal byte ------------")

	byt := []byte(`{"num":6.13, "strs":["a","b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	fmt.Println("---------- Unmarshal response2 ------------")
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// Encoder
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "letture": 7}
	enc.Encode(d)
}
