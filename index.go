package main

import "fmt"

type user struct {
	Name string
	Age  int
}

func main() {

	a := 0
	switch a {
	case 0:
		fmt.Println("a是0")
	case 1:
		fmt.Println("a是1")
	default:
		fmt.Println("a不是0也不是1")
	}
	if a == 1 {
		fmt.Println("a是1")
	}
	if a == 0 {
		fmt.Println("a是0")
	}
	arr := []int{1, 2, 3}
	for key, v := range arr { //range是和JavaScript里面的for ... of 循环类似的
		fmt.Println(key, v)
	}
	add1(1, 2)
	fmt.Println(add1(1, 2))
	var p *int
	p = &a
	*p = 100
	fmt.Println(a) //这里是因为p和a使用了同一个内存了，即共享内存了
	arr1 := make([]int, 3)
	for key, v := range arr1 {
		fmt.Println(key, v)
	}
	map1 := make(map[int]string)
	map1[1] = "张三"
	map1[2] = "李四"
	for key, value := range map1 {
		fmt.Println(key, value)
	}
	us := use{
		Name: "张三",
		Age:  18,
	}
	us.Speak()
	var person Speaker
	person = us
	person.Speak()
}
func add1(a int, b int) int {
	return a + b
}
func sum(a ...int) int {
	s := 0
	for _, v := range a {
		s += v
	}
	return s
}
