package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	c := 100
	return c
}

// 返回多个返回值, 返回值是匿名的
func foo2(a int, b int) (int, int) {
	c := a + b
	d := a - b
	return c, d
}

// 返回多个返回值, 有形参名称的
func foo3(a int, b int) (c int, d int) {
	c = a + b
	d = a - b
	return
}

// 返回值是同一类型的
func foo4(a int, b int) (c, d int) {
	c = a + b
	d = a - b
	return
}

func main() {
	c := foo1("abc", 123)
	fmt.Println(c)
	x, y := foo2(2, 3)
	fmt.Println("- : ", y, ", + : ", x)
	res1, res2 := foo3(5, 4)
	fmt.Println("rse1: ", res1, ", res2: ", res2)
	res3, res4 := foo3(4, 3)
	fmt.Println("rse3: ", res3, ", res4: ", res4)
}
