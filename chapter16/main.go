package main

import "fmt"

func main() {
	var a string
	// pair<statictype:string, value:"abc">
	a = "abc"

	// pair<type:string, value:"abc">
	var allType interface{}
	allType = a

	str, _:=allType.(string)
	fmt.Println(str)
}
