package main

import "fmt"

type Person interface {
	SayHello()
}

type person struct {
	Name     string
	Age      int
	Eyecount int
}

func (p *person) SayHello() {
	fmt.Printf("hello, my name is %s, i have %d years old and %d eyes\n", p.Name, p.Age, p.Eyecount)
}

type tiredperson struct {
	Name     string
	Age      int
	Eyecount int
}

func (tp *tiredperson) SayHello() {
	println("sorry, i am too tired to talk to you right now")
}

func NewPerson(name string, age int) Person {
	if age > 90 {
		return &tiredperson{
			Name:     name,
			Age:      age,
			Eyecount: 2,
		}		
	}
	return &person{
		Name:     name,
		Age:      age,
		Eyecount: 2,
	}
}

func main() {
	p := NewPerson("Rodrigo", 47)
	p.SayHello()

	tp := NewPerson("Baliza", 99)
	tp.SayHello()
}
