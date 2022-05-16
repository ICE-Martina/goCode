package main

import "fmt"

// 定义一个接口, 本质上是一个指针
type Animal interface {
	Sleep()
	GetColor() string
	GetVariety() string
}

type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("cat sleepping...")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetVariety() string {
	return "Cat"
}

type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("dog sleepping...")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetVariety() string {
	return "Dog"
}

func Show(animal Animal) {
	animal.Sleep()
	fmt.Println("cololr: ", animal.GetColor())
	fmt.Println("kind: ", animal.GetVariety())
}

// 多态
// 有一个父类(接口)
// 子类实现父类的全部接口
// 父类类型的变量(指针)指向子类的具体数据变量
func main() {
	// 第一种创建方式：
	var animal Animal
	animal = &Cat{"Yellow"}
	animal.Sleep()

	animal = &Dog{"Black"}
	animal.Sleep()

	fmt.Println("===================")

	// 第二种创建方式：
	cat := Cat{"white"}
	dog := Dog{"Yellow"}
	Show(&cat)
	Show(&dog)
}
