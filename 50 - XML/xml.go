package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMlName xml.Name `xml:plant`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println("-----------")
	fmt.Println(string(out))
	fmt.Println("-----------")
	fmt.Println(xml.Header + string(out))
	fmt.Println("-----------")

	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)
	fmt.Println("-----------")

	tomato := &Plant{Id: 81, Name: "Tomoto"}
	tomato.Origin = []string{"mexico", "California"}

	type Nesting struct {
		XMLName xml.Name `xml:"neting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}
	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
	fmt.Println("-----------")
}
