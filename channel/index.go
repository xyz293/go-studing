package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mutex sync.Mutex

func producer(n int) {
	for a := (n - 1) * 100; a < n*100; a++ {
		if a%5 == 0 {
			fmt.Println(a)
		}
	}
	wg.Done()

}
func read(a chan int, n int) {
	for i := 1; i < n; i++ {
		a <- i
	}
	close(a)
	wg.Done()
}

func main() {
	/*for a := 1; a < 4; a++ {
		wg.Add(1)
		go producer(a)
	}
	wg.Wait()
	fmt.Println("所有任务完成")
	ch := make(chan int, 3)
	ch <- 2
	ch <- 3
	ch <- 4
	close(ch)
	for a := range ch {
		fmt.Println(a)
	}*/
	ch := make(chan int, 5)
	st := make(chan string, 5)
	w := make(chan bool)
	go intworker(4, ch)
	go strworker(4, st)
	go print(ch, w)
	go print1(st, w)
	<-w
	<-w

}
