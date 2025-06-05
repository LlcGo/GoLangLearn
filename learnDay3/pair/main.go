package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Write interface {
	WriteBook()
}

type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("read book")
}

func (this *Book) WriteBook() {
	fmt.Println("write book")
}

func main() {
	// Pair<type:Book value:book地址>
	book := &Book{}

	var r Reader
	// Pair<type:Book value:book地址>
	r = book
	r.ReadBook()

	var w Write
	w = r.(Write)
	// Pair<type:Book value:book地址>
	w.WriteBook()
}
