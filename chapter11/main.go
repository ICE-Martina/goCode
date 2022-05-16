package main

import "fmt"

// 声明一种数据类型的方式
// type myInt int

// 定义一个结构体
type Person struct {
	id   int
	name string
	age  int
}

// 值传递
func passByValue(person Person) {
	person.age = 22
}

// 指针传递
func passByReference(person *Person) {
	person.age = 23
}

func main() {
	// var a myInt = 11
	// fmt.Printf("a type: %T, a=%v\n", a, a)

	var person1 Person

	person1.id = 1
	person1.name = "Jack"
	person1.age = 20
	fmt.Printf("%v\n", person1)

	passByValue(person1)
	fmt.Printf("%v\n", person1)
	
	passByReference(&person1)
	fmt.Printf("%v\n", person1)
}
