package main

import "fmt"

type HunMan1 struct {
	name string
	sex  string
}

func (this *HunMan1) Eat() {
	fmt.Println("HunMan.Eat() ...")
}

func (this *HunMan1) Walk() {
	fmt.Println("HunMan.Walk() ...")
}

type SuperMan1 struct {
	HunMan1
	level int
}

func (this *SuperMan1) Eat() {
	fmt.Println("SuperMan.Eat() ...")
}

func (this *SuperMan1) Fly() {
	fmt.Println("SuperMan.Fly() ...")
}

func main() {
	hunMan1 := HunMan1{"lc", "男"}
	hunMan1.Eat()
	hunMan1.Walk()

	superMan1 := SuperMan1{hunMan1, 1}
	superMan1.Eat()
	superMan1.Walk()
	superMan1.Fly()
}
