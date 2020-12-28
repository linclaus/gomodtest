package model

import "fmt"

type Animal struct {
	Kind string
	Sex  string
}

type People struct {
	Animal
	Name string
}

func TestType() {
	p := People{
		Name: "Mr Lin",
	}
	var t People
	t = p
	t.Name = "Mr Qiu"
	fmt.Printf("p.name:%s\n", p.Name)
	fmt.Printf("t.name:%s\n", t.Name)
	fmt.Printf("sex:%s\n", p.Sex)
}
