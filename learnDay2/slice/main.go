package main

import "fmt"

func main() {
	myArray0 := [10]int{1, 2, 3}
	var myArray1 [10]int

	for index, value := range myArray0 {
		fmt.Println("myArray0 index = ", index, "value =", value)
	}

	for i := 0; i < len(myArray1); i++ {
		fmt.Println("myArray 1 index = ", i, "value = ", myArray1[i])
	}

	fmt.Printf("myArray0 Type %T\n", myArray0)
	fmt.Printf("myArray1 Type %T\n", myArray1)
}
