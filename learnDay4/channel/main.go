package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine 结束")
		for i := 0; i < 3; i++ {
			fmt.Println("管道传值 = ", i)
			// 没人取值阻塞
			c <- i
		}
	}()

	for {
	}

	for i := 0; i < 3; i++ {
		i := <-c
		fmt.Println("管道中获取值 = ", i)
	}
	fmt.Println("结束 ...")
}
