package main

import (
	"fmt"
	"reflect"
)

func reflectNum(arg interface{}) {
	fmt.Println("type: ", reflect.TypeOf(arg))
	fmt.Println("value: ", reflect.ValueOf(arg))
}

type User struct {
	id   int
	name string
	age  int
}

func (this User) Call() {
	fmt.Println("user.Call()...")
	fmt.Printf("%v\n", this)
}

func DoFiledAndMethod(input interface{}) {
	inpuType := reflect.TypeOf(input)
	fmt.Println("type: ", inpuType)

	inputValue := reflect.ValueOf(input)
	fmt.Println("value: ", inputValue)

	for i := 0; i < inpuType.NumField(); i++ {
		field := inpuType.Field(i)
		value := inputValue.Field(i)
		fmt.Printf("%s: %v=%v\n", field.Name, field.Type, value)
	}

	for i := 0; i < inpuType.NumMethod(); i++ {
		method := inpuType.Method(i)
		fmt.Printf("%s: %v\n", method.Name, method.Type)
	}
}

func main() {
	var num float64 = 3.1415
	reflectNum(num)

	user := User{1, "alice", 20}
	DoFiledAndMethod(user)
}
