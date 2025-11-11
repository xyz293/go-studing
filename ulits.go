package main

import "fmt"

type use struct {
	Name string
	Age  int
}
type Speaker interface {
	Speak()
}

func (u use) Speak() {
	fmt.Println("我是", u.Name, "我今年", u.Age, "岁")
}
