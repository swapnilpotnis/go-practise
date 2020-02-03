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
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func (e Employee) saysomething() {
	fmt.Printf("Hi, I am %s and I work at %s Company\n", e.name, e.company)
}

func main() {
	richard := Employee{Human{"Richard", 22, "9999999999"}, "Abc Pvt. Ltd"}
	george := Student{Human{"George", 29, "8888888888"}, "MIT College"}

	richard.saysomething()
	george.saysomething()

}
