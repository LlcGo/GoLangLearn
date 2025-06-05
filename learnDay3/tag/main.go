package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"我的名称"`
	Sex  string `info:"sex"`
}

func findTage(any interface{}) {
	elem := reflect.TypeOf(any).Elem()
	for i := 0; i < elem.NumField(); i++ {
		infoValue := elem.Field(i).Tag.Get("info")
		docValue := elem.Field(i).Tag.Get("doc")
		fmt.Println("info = ", infoValue)
		fmt.Println("doc = ", docValue)
	}
}

func main() {
	re := &resume{}
	findTage(re)
}
