package main

import "fmt"

func main() {
	// slice的声明方式
	// 方式一：
	slice1 := []int{1, 2, 3}
	fmt.Printf("len=%d, slice1=%v\n", len(slice1), slice1)

	// 方式二：声明一个slice
	var slice2 []int
	fmt.Printf("len=%d, slice2=%v\n", len(slice2), slice2)
	if slice2 == nil {
		fmt.Println("slice2 没有分配空间")
	}
	// 分配空间
	slice2 = make([]int, 3)
	slice2[2] = 1
	fmt.Printf("len=%d, slice2=%v\n", len(slice2), slice2)

	// 方式三：
	var slice3 []int = make([]int, 3)
	fmt.Printf("len=%d, slice3=%v\n", len(slice3), slice3)

	// 方式四：
	slice4 := make([]int, 3)
	fmt.Printf("len=%d, slice4=%v\n", len(slice4), slice4)
}
