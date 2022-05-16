package main

import "fmt"

// 继承
type Person struct {
	name   string
	gender string
}

func (this *Person) Eat() {
	fmt.Println("person.Eat()...")
}

func (this *Person) Walk() {
	fmt.Println("person.Walk()...")
}

type SuperMan struct {
	Person // SuperMan继承了Person

	level int
}

// 重定义父类的方法
func (this *SuperMan) Eat(){
	fmt.Println("SuperMan.Eat()...")
}

// 子类的新方法
func (this *SuperMan) Fly(){
	fmt.Println("SuperMan.Fly()...")
}

func (this * SuperMan) Show(){
	fmt.Println("name=", this.name)
	fmt.Println("gender=", this.gender)
	fmt.Println("level=", this.level)
}

func main() {
	p := Person{"Bobu", "male"}
	p.Eat()
	p.Walk()

	// 第一种创建子类的方式：
	s :=SuperMan{Person{"Jack", "male"}, 70}
	s.Eat()
	s.Walk()
	s.Fly()

	// 第二种创建子类的方式：
	var s1 SuperMan
	s1.name = "Lucy"
	s1.gender = "female"
	s1.level = 90

	s1.Show()
}
