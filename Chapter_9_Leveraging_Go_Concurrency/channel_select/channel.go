package main

import (
	"fmt"
)

func callerA(c chan string) {
	c <- "Hello World!"
	close(c)
}

func callerB(c chan string) {
	c <- "Hola Mundo!"
	close(c)
}

func main() {
	a, b := make(chan string), make(chan string)
	go callerA(a)
	go callerB(b)

	var msg string
	ok1, ok2 := true, true

	for ok1 || ok2 {
		select {
		case msg, ok1 = <-a:
			if ok1 {
				fmt.Printf("A: %s\n", msg)
			}
		case msg, ok2 = <-b:
			if ok2 {
				fmt.Printf("B: %s\n", msg)
			}
		}
	}

	//for i := 0; i < 5; i++ {
	//	// 需要等待通道 a 和 b，已经接收到值了
	//	time.Sleep(1 * time.Microsecond)
	//	select {
	//	case msg := <-a:
	//		fmt.Printf("A: %s\n", msg)
	//	case msg := <-b:
	//		fmt.Printf("B: %s\n", msg)
	//	default:
	//		fmt.Println("No message received")
	//	}
	//}
}
