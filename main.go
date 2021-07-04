package main

import (
	"fmt"

	"github.com/Jinwook-Song/nomadcoin-2021/person"
)


func main() {
	nico := person.Person{}
	nico.SetDetails("jinwook", 29)
	fmt.Println("Main 'person", nico)
	fmt.Println(nico.Name())
}

