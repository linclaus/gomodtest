package test

import (
	"testing"

	"github.com/sirupsen/logrus"
)

type notifyer interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u user) notify() {
	logrus.Printf("name: %s, email: %s\n", u.name, u.email)
}

type admin struct {
	name string
	age  int
}

func (a admin) notify() {
	logrus.Printf("name: %s, age: %d\n", a.name, a.age)
}

func sendNotify(n notifyer) {
	n.notify()
}

func TestPolymorphism(t *testing.T) {
	u := user{
		name:  "user",
		email: "email",
	}
	a := admin{
		name: "admin",
		age:  10,
	}
	sendNotify(u)
	sendNotify(a)
}
