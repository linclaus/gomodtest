package test

import (
	"fmt"
	"reflect" // 这里引入reflect模块

	"gopkg.in/yaml.v2"
)

type MPeople struct {
	Sex string `yaml:"sex,omitempty"`
	Age int64  `yaml:"age,omitempty"`
}

type MStudent struct {
	People `yaml:",inline"`
	Class  string `yaml:"class,omitempty"`
}

type User struct {
	Name   string "user name" //这引号里面的就是tag
	Passwd string "user passsword"
}

func TestStructTag() {
	data := `
sex: man
age: 14
class: one
`
	var b MStudent
	yaml.Unmarshal([]byte(data), &b)
	fmt.Println(b)

}

func Test() {
	user := &User{"chronos", "pass"}
	s := reflect.TypeOf(user).Elem() //通过反射获取type定义
	for i := 0; i < s.NumField(); i++ {
		fmt.Println(s.Field(i).Tag) //将tag输出出来
	}
}
