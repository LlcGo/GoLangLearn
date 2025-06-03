package main

import "fmt"

func printMap(myMap map[string]string) {
	for key, value := range myMap {
		fmt.Println("key =", key)
		fmt.Println("value =", value)
	}
}

func changeValue(myMap map[string]string) {
	myMap["a"] = "e"
}

func main() {
	// 添加
	myMap := make(map[string]string)
	myMap["a"] = "b"
	myMap["c"] = "d"
	myMap["e"] = "f"

	// 遍历
	printMap(myMap)

	// 删除
	fmt.Println("*****************删除后*****************")
	delete(myMap, "a")
	printMap(myMap)

	// 修改map
	fmt.Println("*****************修改后*****************")
	changeValue(myMap)
	printMap(myMap)
}
