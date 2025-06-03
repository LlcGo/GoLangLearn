package main

import (
	"fmt"
	"strconv"
)

func main() {
	var myMap0 map[string]string
	if myMap0 == nil {
		fmt.Println("map为空")
	}

	myMap1 := make(map[string]string, 10)
	myMap1["aaa"] = "bbb"
	myMap1["bbb"] = "ccc"
	myMap1["ddd"] = "eee"
	fmt.Println(myMap1)
	for s := range myMap1 {
		fmt.Println(s)
		fmt.Println(myMap1[s])
	}

	myMap2 := make(map[int]string)
	fmt.Println(len(myMap2))
	for s := range myMap2 {
		fmt.Println(s)
		fmt.Println(myMap1[strconv.Itoa(s)])
	}

	myMap3 := map[string]string{
		"aaaa": "bbbb",
		"cccc": "eeee",
		"dddd": "bbbb",
	}
	fmt.Println(myMap3)
}
