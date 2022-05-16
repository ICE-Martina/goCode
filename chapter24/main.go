package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		close(c)
	}()

	/* 	for {
		if data, ok := <-c; ok {
			fmt.Println(data)
		}else{
			break
		}
	} */

	// 可以使用range来不断迭代channel
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("main finish !")
}
