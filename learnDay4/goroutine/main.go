package main

import (
	"fmt"
	"time"
)

func PrintFor() {
	for i := 0; i < 10; i++ {
		fmt.Println("go i = ", i)
	}
}

func main() {
	go PrintFor()

	for i := 0; i < 10; i++ {
		fmt.Println("i = ", i)
	}

	time.Sleep(1)
}
