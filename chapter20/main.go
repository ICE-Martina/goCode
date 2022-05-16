package main

import (
	"fmt"
	_ "runtime"
	"time"
)

// 子goroutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine: i=%d\n", i)
		time.Sleep(1 * time.Second)
	}
}

// 主goroutine
func main() {

	// main线程执行newTask()
	// newTask()

	// 创建一个go程执行newTask()
	// go newTask()

	// 匿名函数，子goroutine
	go func() {
		defer fmt.Println("A.defer...")
		// return // 退出go程
		func() {
			defer fmt.Println("B.defer...")
			// runtime.Goexit() // 退出go程
			fmt.Println("B...")
		}()
		fmt.Println("A...")
	}()

	go func(a int, b int) bool {
		fmt.Println("a=", a, " b=", b)
		return a > b

	}(20, 10)

	// 死循环
	for {
		time.Sleep(1 * time.Second)
	}
}
