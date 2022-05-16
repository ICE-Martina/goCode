package main

import "fmt"

// goroutine之间的通信
func main() {
	// 定义channel(无缓冲区的channel)
	c := make(chan int)

	go func() {
		defer fmt.Println("child goroutine finish!")
		fmt.Println("child goroutine exec...")
		c <- 666 // 将666发送给c
	}()

	// <-c        // 只是接收c中数据，并丢弃
	// num := <-c // 从c中接收数据，并赋值给num
	// fmt.Println("num=", num)

	num1, ok := <-c // 从c中接收数据，赋值给num1，并返回接收数据是否有异常ok

	if ok {
		fmt.Println("num1=", num1)
	}
	fmt.Println("main goroutine finish")
}
