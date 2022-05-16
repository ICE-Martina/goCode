package main

import "fmt"

// 值传递, 交换不成功
func swap(b int, c int) {
	var temp int
	temp = b
	b = c
	c = temp
}

func swap1(b *int, c *int) {
	var temp int
	temp = *b
	*b = *c
	*c = temp
}

func main() {
	var a int = 1
	// 值传递
	changeValue1(a)
	fmt.Println("a = ", a)

	// 地址传递(引用传递)
	changeValue2(&a)
	fmt.Println("a = ", a)

	var b, c int = 3, 7
	swap(b, c)
	fmt.Println("b=", b, ", c=", c)

	swap1(&b, &c)
	fmt.Println("b=", b, ", c=", c)
}

func changeValue1(p int) {
	p = 10
}

func changeValue2(p *int) {
	*p = 10
}
