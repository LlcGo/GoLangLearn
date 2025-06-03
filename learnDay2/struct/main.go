package main

import "fmt"

type myInt int

type Book struct {
	auth  string
	title string
}

func printBook(book *Book) {
	fmt.Printf("book.auth = %s , boo.title = %s\n", book.auth, book.title)
}

func changeBook0(book Book) {
	book.auth = "lc1"
}

func changeBook1(book *Book) {
	book.auth = "lc1"
}

func main() {
	var testInt myInt
	fmt.Printf("type = %T\n", testInt)

	var book Book
	book.auth = "lc"
	book.title = "go"
	fmt.Println("book = ", book)

	changeBook0(book)
	printBook(&book)

	changeBook1(&book)
	printBook(&book)
}
