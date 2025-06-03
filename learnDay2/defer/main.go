package main

import "fmt"

func deferFunc1() {
	fmt.Println("def 1 ...")
}

func deferFunc2() {
	fmt.Println("def 2 ...")
}

func deferFunc3() {
	fmt.Println("def 3 ...")
}

func returnInt() int {
	fmt.Println("return ...")
	return 0
}

func function() int {
	defer deferFunc1() // 后执行
	return returnInt() // 先执行
}

func main() {
	// defer最后执行 并且压栈 后进先执行 3 2 1
	defer deferFunc1()
	defer deferFunc2()
	defer deferFunc3()

	function()
}
