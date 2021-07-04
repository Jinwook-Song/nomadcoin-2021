package main

import "fmt"

type person struct {
	name string
	age int
}

// method
func (p person) sayHello() {
	fmt.Printf("Hello! My name is %s and I'm %d", p.name, p.age)
}

func main() {
	nico := person{"jinwook", 29}
	nico.sayHello()
}