package main

import (
	"fmt"
	"reflect"
)

type person struct {
	Name   string
	Age    int
	height float64
}

func structfile(x interface{}) {
	v := reflect.TypeOf(x)
	for i := 0; i < v.NumField(); i++ {
		file := v.Field(i) //这里是获取每一个结构体里面的字段的
		fmt.Println(file.Name, file.Type)
	}
}
func fn(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Println(v)
	s := reflect.ValueOf(x)
	if s.Kind() == reflect.Func {
		arg := []reflect.Value{ //这步是传入的参数
			reflect.ValueOf(100),
			reflect.ValueOf(200)}
		res := s.Call(arg)
		fmt.Println(res[0])
	}
}
func get(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Struct {
		s := v.FieldByName("Name")
		fmt.Println(s)
		a := v.FieldByName("Age")
		fmt.Println(a)
	}
}
func add(a int, b int) []int {
	return []int{a , b}
}
func main() {
	s := person{
		Name:   "张三",
		Age:    18,
		height: 1.8,
	}
	//structfile(s)
	get(s)
	fn(add)

}
