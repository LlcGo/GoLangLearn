package main

import "fmt"

type SuperMan struct {
	HunMan
	level int
}

func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat() ...")
}

func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly() ...")
}
