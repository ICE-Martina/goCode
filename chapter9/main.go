package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 1)
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 2)
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 3)
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(numbers), cap(numbers), numbers)

	numbers1 := make([]int, 2)
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(numbers1), cap(numbers1), numbers1)

	numbers1 = append(numbers1, 1)
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(numbers1), cap(numbers1), numbers1)

	num2 := []int{1, 2, 3, 4}

	// [0,2)
	num3 := num2[0:2]
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(num3), cap(num3), num3)
	num2[0] = 11
	// num2和num3指向同内存地址
	fmt.Println(num2)
	fmt.Println(num3)

	num4 := make([]int, 4)
	// num4和num2指向不同内存地址
	copy(num4, num2)
	fmt.Println(num4)
}
