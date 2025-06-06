package main

import "fmt"

func main() {
	c := make(chan int, 3)
	go func() {
		for i := 0; i <= 5; i++ {
			fmt.Println("管道传值 = ", i)
			c <- i
		}
	}()

	for i := 0; i < 3; i++ {
		number := <-c
		fmt.Println("管道获取值 = ", number)
	}

	//for {
	//	number := <-c
	//	fmt.Println("管道获取值 = ", number)
	//}

	for {
	}
}
