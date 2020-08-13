package test

import "fmt"

type People struct {
	Sex string
	Age int64
}

type Student struct {
	People
	Class string
}

type School struct {
	Students []People
	Name     string
}

func TestExtend() {

	p := People{
		Sex: "man",
		Age: 16,
	}
	st := Student{
		People: p,
		Class:  "class10",
	}

	sc := School{
		Students: []People{p},
		Name:     "school1",
	}
	fmt.Printf("people: %s\n", p)
	fmt.Printf("Student: %s\n", st)
	fmt.Printf("School: %s\n", sc)
}
