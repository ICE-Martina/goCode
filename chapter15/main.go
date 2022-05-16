package main

import "fmt"

// 通用类型->空接口
// interface{}是万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc execute...")
	fmt.Println(arg)

	// interface{}的类型断言机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg 不是string类型")
	} else {
		fmt.Printf("value type: %T\n", value)
		fmt.Println("arg是string类型, value=", value)
	}
}

type Book struct {
	author string
}

func main() {
	book := Book{"ZS"}
	myFunc(book)
	myFunc(111)
	myFunc("abc")
}
