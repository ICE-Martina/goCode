package main

import (
	"fmt"
	"time"
)

func main() {
	// 定义一个有缓冲区的channel
	// 当channel已经满，再向channel写数据就会阻塞
	// 当channel为空，再从channel读数据旧会阻塞
	c := make(chan int, 3)
	fmt.Println("c-> len(c):", len(c), "cap(c):", cap(c))

	go func() {
		defer fmt.Println("子go程结束!")
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("子go程写入channel的数据: ", i, "channel的len: ", len(c), "channel的cap: ", cap(c))
		}

	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println("num=", num, "len(c)", len(c))
	}

	time.Sleep(1 * time.Second)
	fmt.Println("main goroutine finish !")
}
