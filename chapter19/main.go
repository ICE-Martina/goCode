package main

import (
	"encoding/json"
	"fmt"
)

// tag的使用
type Movie struct {
	Title string   `json:"title"`
	Year  int      `json:"year"`
	Price int      `json:"price"`
	Actor []string `json:"actors"`
}

func main() {

	movie := Movie{"整蛊专家", 1997, 30, []string{"zhouxingchi", "wumengda"}}

	// 编码过程：结构体->json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("struct to json error", err)
		return
	}
	fmt.Printf("result= %s\n", jsonStr)

	// 解码过程 json->结构体
	toMovie := Movie{}
	err = json.Unmarshal(jsonStr, &toMovie)
	if err != nil {
		fmt.Println("json to struct error", err)
		return
	}
	fmt.Printf("%v\n", toMovie)
}
