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
	fmt.Printf("kind:%s\n", p.Kind)
	fmt.Printf("sex:%s\n", p.Sex)
}
