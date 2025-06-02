package main

import "fmt"

var e int
var f int = 100
var g = 100

// d: = 0.99 无法全局声明

func main() {
	// 方法一:声明变量,int默认值为0
	var a int
	fmt.Println("a = ", a)
	// 方法二:声明一个变量,初始化一个值
	var b int = 100
	fmt.Println("b = ", b)
	// 方法三:初始化省略数据类型,默认识别
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("c Type = %T\n", c)
	// 方法四:省去var(常用) 但是无法全局声明
	d := 0.99
	fmt.Println("d = ", d)
	fmt.Printf("d Type = %T\n", d)

	// 声明多个变量
	var aa, bb = 100, 100
	fmt.Println("aa = ", aa, "bb = ", bb)
	// 多行变量
	var (
		cc = 100
		dd = 100
	)
	fmt.Println("cc = ", cc, "dd = ", dd)
}
