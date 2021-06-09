package test

import (
	"fmt"
	"reflect" // 这里引入reflect模块
	"testing"

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

func TestStructTag(t *testing.T) {
	data := `
sex: man
age: 14
class: one
`
	var b MStudent
	yaml.Unmarshal([]byte(data), &b)
	fmt.Println(b)

}

func Test(t *testing.T) {
	user := &MPeople{"chronos", 11}
	s := reflect.TypeOf(user).Elem() //通过反射获取type定义
	for i := 0; i < s.NumField(); i++ {
		fmt.Println(s.Field(i).Tag) //将tag输出出来
	}
}
