package main

import (
	"fmt"
)

// 声明全局变量时，方法一、方法二、方法三是可以的
var v int
var aa = "!!!!"
var bb string = "####"

// 方法四只能在函数体内来声明

func main() {
	/**
	* 四种变量的声明方式
	 */

	// 方法一：
	var a int
	fmt.Println(a)

	// 方法二：
	var b int = 1
	fmt.Println(b)

	// 方法三：
	var c = 2
	fmt.Println(c)
	fmt.Printf("c type: %T\n", c)

	// 方法四：
	d := 3
	fmt.Println(d)
	fmt.Printf("d type: %T\n", d)

	fmt.Println("aa: ", aa, ", bb: ", bb)

	var xx, yy int = 11, 22
	fmt.Println("xx: ", xx, ", yy: ", yy)

	var cc, dd = 2.1, "111"
	fmt.Println("cc: ", cc, ", dd", dd)

	var (
		aaa int  = 100
		bbb bool = true
	)
	fmt.Println("aaa: ", aaa, ", bbb: ", bbb)
}
