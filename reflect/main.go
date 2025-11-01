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
	fmt.Println(v)
}

func main() {
	reflectFn(100)
	reflectFn("hello")
	reflectFn(user{
		Name: "张三",
		Age:  18,
	})
}
