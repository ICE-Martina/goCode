package main

import (
	"fmt"
)

// map指针传递(引用传递)
func mapFunc(args map[int]string) {
	args[5] = "wuhan"
	for key, value := range args {
		fmt.Println("key->", key, "value->", value)
	}
}

func main() {
	/*
		// 声明方式一：
		var map1 map[string]string
		map1 = make(map[string]string, 3)

		map1["a"] = "1"
		map1["b"] = "2"
		map1["c"] = "3"

		fmt.Println(map1)

		// 声明方式二：
		map2 := make(map[int]string)
		fmt.Printf("len=%d, map2=%v\n", len(map2), map2)
		map2[1] = "a"
		fmt.Printf("len=%d, map2=%v\n", len(map2), map2)
		map2[2] = "aa"
		map2[3] = "aaa"
		fmt.Println(map2)

		// 声明方式三：
		map3 := map[int]string{
			1: "1",
			2: "11",
			3: "111",
		}
		fmt.Println(map3)
	*/

	city := make(map[int]string)

	city[1] = "beijing"
	city[2] = "shanghai"
	city[3] = "shenzheng"
	for key, value := range city {
		fmt.Println("key=", key, ", value=", value)
	}

	delete(city, 2)

	city[1] = "chengdu"

	fmt.Println("=====================================")
	for key, value := range city {
		fmt.Println("key=", key, ", value=", value)
	}

	mapFunc(city)

}
