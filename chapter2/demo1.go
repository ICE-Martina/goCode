package main

import "fmt"

// const 来定义枚举类型
const (
	// 可以在const()添加一个关键字iota，每行的iota都会累加1，第一行的iota默认值是0, iota只能在const()中使用
	// BEIJING   = 1
	// SHANGHAI  = 2
	// SHENZHENG = 3
	BEIJING = iota
	SHANGHAI
	SHENZHENG
)

func main() {
	var a int = 1
	fmt.Println(a)
	// 常量(只读)
	const b int = 2
	fmt.Println(b)

	fmt.Println("BEIJING: ", BEIJING)
	fmt.Println("SHANGHAI: ", SHANGHAI)
	fmt.Println("SHENZHENG: ", SHENZHENG)
}
