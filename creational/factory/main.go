package main

import "fmt"

type Person struct {
	Name     string
	Age      int
	Eyecount int
}

func NewPerson1(name string, age int) Person {
	return Person{name, age, 2}
}

func NewPerson2(name string, age int) *Person {
	return &Person{name, age, 2}
}

func main_() {
	// initialize directly
	p := Person{"John", 22, 2}
	fmt.Println(p)

	// use a constructor
	p1 := NewPerson1("Mary", 20)
	fmt.Println(p1)

	// use a constructor
	p2 := NewPerson2("Jane", 20)
	fmt.Println(p2)
}
