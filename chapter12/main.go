package main

import "fmt"

// 封装
// 如果类型首字母大写, 表示其他包也可以访问
type Hero struct {
	name  string
	// 如果属性首字母大写, 表示该属性对外可以访问, 否则只能类内部访问
	Ad    int
	level int
}

// 值传递
/*
func (this Hero) getName() string {
	return this.name
}

func (this Hero) setName(name string) {
	this.name = name
}

func (this Hero) show() {
	fmt.Println("name = ", this.name)
	fmt.Println("ad = ", this.ad)
	fmt.Println("level = ", this.level)
}
*/

// 指针传递
func (this *Hero) getName() string {
	return this.name
}

func (this *Hero) setName(name string) {
	this.name = name
}

func (this *Hero) show() {
	fmt.Println("name = ", this.name)
	fmt.Println("ad = ", this.Ad)
	fmt.Println("level = ", this.level)
}

func main() {

	hero := Hero{name: "jack", Ad: 50, level: 2}
	hero.show()
	hero.setName("Lucy")
	hero.show()

}
