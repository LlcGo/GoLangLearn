package main

import "fmt"

// 大写代表公开 小写代表私有其他包不可访问
type Person struct {
	Name string
	// 内部属性也是
	age int
}

func (this *Person) GetName() string {
	return this.Name
}

func (this *Person) SetName(name string) {
	this.Name = name
}

func (this Person) Show() {
	fmt.Println("person = ", this)
}

func main() {
	person := Person{
		Name: "lc",
		age:  25,
	}
	person.Show()

	name := person.GetName()
	fmt.Println("name = ", name)

	person.SetName("lc1")
	person.Show()
}
