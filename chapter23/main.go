package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {

		for i := 0; i < 4; i++ {
			c <- i
		}

		// 如果不关闭channel会报错: fatal error: all goroutines are asleep - deadlock!
		// 关闭channel
		close(c)
	}()

	for {
		// ok如果为true表示channel没有关闭，为false表示已经关闭
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Println("main finish !")
}
