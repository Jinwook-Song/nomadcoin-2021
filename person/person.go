package person

import "fmt"

// uppercase
type Person struct {
	name string
	age int
}

// receiver pointer function
// mutate struct not just copy thing
func (p *Person) SetDetails(name string, age int) {
	p.name = name
	p.age = age
	fmt.Println("SeeDetail's person", p)
}

func (p Person) Name() string {
	return p.name
}