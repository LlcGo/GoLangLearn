package main

import "fmt"

type HunMan struct {
	name string
	sex  string
}

func (this *HunMan) Eat() {
	fmt.Println("HunMan.Eat() ...")
}

func (this *HunMan) Walk() {
	fmt.Println("HunMan.Walk() ...")
}
