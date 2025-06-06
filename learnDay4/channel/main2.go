package main

import "fmt"

func main() {
	c := make(chan int, 2)

	go func(number1 int, number2 int) {
		for i := 0; i < 5; i++ {
			c <- number1 + number2
		}
		close(c)
	}(1, 2)

	for {
		if data, ok := <-c; ok {
			fmt.Println("c = ", data)
		} else {
			break
		}
	}

	//for i := range c {
	//	number := <-c
	//	fmt.Println("i = ", i, "number = ", number)
	//}

}
