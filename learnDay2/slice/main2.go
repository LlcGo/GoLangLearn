package main

import "fmt"

func main() {
	// 方法1
	slice0 := []int{1, 2, 3}
	// 方法2
	var slice1 []int
	// 方法3
	var slice2 = make([]int, 3)
	// 方法四
	slice3 := make([]int, 3)

	//打印
	fmt.Printf("slice0 len = %d vlaue = %v\n", len(slice0), slice0)
	fmt.Printf("slice1 len = %d vlaue = %v\n", len(slice1), slice1)
	fmt.Printf("slice2 len = %d vlaue = %v\n", len(slice2), slice2)
	fmt.Printf("slice3 len = %d vlaue = %v\n", len(slice3), slice3)
}
