package main

import "fmt"

func foo1(a int, b string) int {
	fmt.Println("a = ", a, "b = ", b)
	c := 100
	return c
}

func foo2(a int, b string) (int, int) {
	fmt.Println("a = ", a, "b = ", b)
	c := 100
	d := 100
	return c, d
}

func foo3(a int, b string) (r1 int, r2 int) {
	fmt.Println("a = ", a, "b = ", b)
	r1 = 100
	r2 = 200
	return
}

func foo4(a int, b string) (r1, r2 int) {
	fmt.Println("a = ", a, "b = ", b)
	r1 = 100
	r2 = 200
	return
}

func main() {
	c := foo1(100, "a")
	fmt.Println("c = ", c)

	r1, r2 := foo2(100, "b")
	fmt.Println("r1 = ", r1, "r2 = ", r2)

	r3, r4 := foo3(100, "c")
	fmt.Println("r3 = ", r3, "r4 = ", r4)

	r5, r6 := foo4(100, "d")
	fmt.Println("r5 = ", r5, "r6 = ", r6)
}
