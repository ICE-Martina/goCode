package main

import "fmt"

// 这种方式传入的固定长度数组，是值传递(把传入值拷贝给形参, 对原来的值不会发生任何影响(例如在函数内改变值))
func arrFunc(arr [4]int) {
	for index, value := range arr {
		fmt.Println("index=", index, ", value=", value)
	}
	arr[2] = 22
}

// 引用传递
func arrFunc1(arr []int) {
	for _, value := range arr {
		fmt.Println(value)
	}
	arr[2] = 222
}

func main() {
	// 1.固定长度数组
	var arr [3]int

	arr1 := [3]int{1, 2}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
		fmt.Println(arr1[i])
	}

	for index, value := range arr1 {
		fmt.Println("index=", index, " value=", value)
	}

	arr2 := [4]int{1, 2, 3, 4}
	arrFunc(arr2)
	fmt.Printf("arr2 type: %T\n", arr2)

	fmt.Println("===============================")
	// 2.切片slice(动态数组)
	arr3 := []int{6, 7, 8, 9}
	fmt.Printf("arr3 type: %T\n", arr3)

	arrFunc1(arr3)
	for _, value := range arr3 {
		fmt.Println(value)
	}
}
