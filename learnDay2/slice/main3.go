package main

import "fmt"

func main() {
	var slice0 = []int{1, 2, 3}
	fmt.Printf("slice0 len = %d, cap = %d, value = %v\n", len(slice0), cap(slice0), slice0)
	slice0 = append(slice0, 100) // 扩容：为原来容量的2倍
	fmt.Printf("slice0 len = %d, cap = %d, value = %v\n", len(slice0), cap(slice0), slice0)

	slice1 := make([]int, 5, 5)
	fmt.Printf("slice1 len = %d, cap = %d, value = %v\n", len(slice1), cap(slice1), slice1)
	slice1 = append(slice1, 100) // 扩容：为原来容量的2倍
	fmt.Printf("slice1 len = %d, cap = %d, value = %v\n", len(slice1), cap(slice1), slice1)

	number0 := slice0[1:3]
	fmt.Printf("number0 len = %d, cap = %d, value = %v\n", len(number0), cap(number0), number0)
	slice3 := make([]int, 5)
	copy(slice3, slice0)
	fmt.Printf("slice3 len = %d, cap = %d, value = %v\n", len(slice3), cap(slice3), slice3)
}
