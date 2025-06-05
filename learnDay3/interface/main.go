package main

import "fmt"

func myFunc(arg interface{}) {
	fmt.Println("arg =", arg)
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg not string")
	} else {
		fmt.Println("arg is string value = ", value)
	}
}

type Book struct {
	title string
}

func main() {
	book := Book{
		title: "Golang",
	}
	myFunc(book)

	myFunc(100)

	myFunc("abc")
}
