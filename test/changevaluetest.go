package test

import "fmt"

type value struct {
	value int
}

func change(v value) {
	v.value = 11
}
func changePtr(v *value) {
	v.value = 11
}

func TestValue() {
	v := value{
		value: 10,
	}
	fmt.Printf("v before change:%d \n", v.value)
	change(v)
	fmt.Printf("v after change:%d \n", v.value)
	t := v
	fmt.Printf("t before change:%d \n", t.value)
	changePtr(&t)
	fmt.Printf("v after change:%d \n", v.value)
	fmt.Printf("t after change:%d \n", t.value)
}
