package main

import "fmt"

/**
 * 单流程下一个go程只能监控一个channel的状态，select可以完成监控多个channel的状态
 */
func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			// 如果c可写，则读case就会进来
			x = y
			y = x + y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	// sub go
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	// main go
	fibonacci(c, quit)
}
