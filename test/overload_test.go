package test

import (
	"testing"

	"github.com/sirupsen/logrus"
)

type Action interface {
	Say()
}

type People struct {
	Sex string
	Age int64
}

func (p People) Say() {
	logrus.Printf("%v", p)
}

type Student struct {
	People
	Class string
}

func (s Student) Say() {
	logrus.Printf("%v", s)
}

func DoAction(a Action) {
	a.Say()
}

type School struct {
	Students []People
	Name     string
}

func TestExtend(t *testing.T) {
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
	logrus.Printf("people: %v\n", p)
	logrus.Printf("Student: %v\n", st)
	logrus.Printf("School: %v\n", sc)
	DoAction(p)
	DoAction(st)
}
