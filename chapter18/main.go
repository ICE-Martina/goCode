package main

import (
	"fmt"
	"reflect"
)

// Tag标签
type resume struct {
	Name   string `info:"name" doc:"field name"`
	gender string `info:"gender" doc:"field gender"`
}

func reflectTag(str interface{}) {
	t := reflect.TypeOf(str)
	for i := 0; i < t.NumField(); i++ {
		info := t.Field(i).Tag.Get("info")
		doc := t.Field(i).Tag.Get("doc")
		fmt.Println("info:", info, ", doc:", doc)

	}
}

func main() {
	var re resume
	reflectTag(re)
}
