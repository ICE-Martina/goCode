package main

import "fmt"

func deferFunc() string {
	fmt.Println("defer exec ...")
	return "defer exec ..."
}

func returnFunc() string {
	fmt.Println("return exec ...")
	return "return exec ..."
}

func deferAndReturn() string {

	defer deferFunc()
	return returnFunc()
}

func main() {

	// 在函数生命周期结束之前执行, 先写的defer先入栈
	defer fmt.Println("********")
	defer fmt.Println("----------")

	fmt.Println("==========")

	deferAndReturn()
}
