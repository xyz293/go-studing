package main

import (
	"fmt"
	"reflect"
)

type user struct {
	Name string
	Age  int
}

func reflectFn(x interface{}) {
	v := reflect.TypeOf(x)
	xv := reflect.ValueOf(x)
	fmt.Println(xv.Kind())
	fmt.Println(v)
}

func set(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		v.SetInt(1000)
	}
}
func main() {
	/*
		reflectFn(100)
		reflectFn("hello")
		reflectFn(user{
			Name: "张三",
			Age:  18,
		})
	*/
	x := 10
	set(&x)
	fmt.Println(x)
}
