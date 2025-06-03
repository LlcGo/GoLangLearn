package main

import "fmt"

// 动态数据引用传递可改变数组内部值
func changeArray(array []int) {
	array[0] = 100
}

func main() {
	myArray := []int{1, 2, 3}
	for i := 0; i < len(myArray); i++ {
		fmt.Println("index = ", i, "value = ", myArray[i])
	}

	changeArray(myArray)
	for i := 0; i < len(myArray); i++ {
		fmt.Println("index = ", i, "value = ", myArray[i])
	}
}
