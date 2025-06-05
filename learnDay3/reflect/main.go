package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func main() {
	user := User{1, "abc", 18}
	DoFieldAndMethod(user)
}

func DoFieldAndMethod(any interface{}) {
	inputValue := reflect.ValueOf(any)
	inputType := reflect.TypeOf(any)

	for i := 0; i < inputValue.NumField(); i++ {
		field := inputValue.Field(i)
		fmt.Println("value field = ", field)
	}

	for i := 0; i < inputType.NumField(); i++ {
		typeField := inputType.Field(i)
		fmt.Println("type = ", typeField.Type,
			"name =", typeField.Name)
	}
}
