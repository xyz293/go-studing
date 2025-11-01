package main

import "fmt"

func intworker(id int, ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- id
	}
	close(ch)
}
func strworker(id int, ch chan string) {
	for i := 0; i < 5; i++ {
		ch <- fmt.Sprintf("工人 %d 开始工作第 %d 次", id, i+1)
	}
	close(ch)

}

func print(ch chan int, w chan bool) {
	for a := range ch {
		fmt.Println(a)
	}
	w <- true
}
func print1(ch chan string, w chan bool) {
	for a := range ch {
		fmt.Println(a)
	}
	w <- true
}
