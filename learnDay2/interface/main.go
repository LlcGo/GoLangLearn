package main

import (
	"fmt"
)

type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {
	Color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat Sleep ...")
}

func (this *Cat) GetColor() string {
	fmt.Println("Cat GetColor ...")
	return this.Color
}

func (this *Cat) GetType() string {
	fmt.Println("Cat GetType ...")
	return "Cat"
}

type Dog struct {
	Color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog Sleep ...")
}

func (this *Dog) GetColor() string {
	fmt.Println("Dog GetColor ...")
	return this.Color
}

func (this *Dog) GetType() string {
	fmt.Println("Dog GetType ...")
	return "Dog"
}

func printAnimal(animal AnimalIF) {
	animal.Sleep()
	animal.GetType()
	animal.GetColor()
}

func main() {
	/**
	var animal AnimalIF

	animal = &Cat{"blue"}
	animal.Sleep()
	animal.GetType()
	animal.GetColor()

	animal = &Dog{"write"}
	animal.Sleep()
	animal.GetType()
	animal.GetColor()
	*/
	cat := Cat{"blue"}
	printAnimal(&cat)
	dog := Dog{"write"}
	printAnimal(&dog)
}
