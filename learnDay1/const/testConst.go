package main

import "fmt"

// 常量枚举
const (
	a = 1
	b = 2
)

// 默认累加
const (
	c = iota
	d
	e
)

const (
	f, g = iota + 1, iota + 2
	h, i

	j, k = iota * 1, iota * 2
	L, m
)

const (
	code1 = 200 + iota
	code2
	code3
)

func main() {
	const name string = "小李"
	fmt.Println("name = ", name)
	// name = "a" 不能修改常量

	fmt.Println("f = ", f, "g = ", g)
	fmt.Println("h = ", h, "i = ", i)
	fmt.Println("j = ", j, "k = ", k)
	fmt.Println("l = ", L, "m = ", m)

	// var a int = iota iota是配合const()一起使用的
}
