package main

import (
	"fmt"
	"time"
)

func thrower(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
		fmt.Printf("throwing >> %d\n", i)
	}
}

func catcher(c chan int) {
	for i := 0; i < 10; i++ {
		num := <-c
		fmt.Printf("catching << %d\n", num)
	}
}

func main() {
	// 无缓冲通道
	//c := make(chan int)
	// 有缓冲通道
	c := make(chan int, 3)
	go thrower(c)
	go catcher(c)
	time.Sleep(100 * time.Millisecond)
}
