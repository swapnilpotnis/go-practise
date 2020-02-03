package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

func (h *Human) saysomething() {
	fmt.Printf("hi...my name is %s and this is a simple function\n", h.name)
}

func main() {
	swap := Employee{Human{"Swap", 22, "9999999999"}, "Abc Company Name"}
	raj := Student{Human{"Raj", 30, "9999999999"}, "Primary School"}

	swap.saysomething()
	raj.saysomething()
}
