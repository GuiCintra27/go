package main

import "fmt"

type Vehicle interface {
	Honk()
}

type Car struct {
	Model string
	Year int
	Manufacturer string
}

type Person struct {
	Name string
	Age int
	Vehicle Car
}

func (c Car) Honk(owner string) {
	fmt.Printf("%s %s is honking\n", owner, c.Model)
}

func (p Person) Drive() {
	fmt.Printf("%s is driving a %s\n", p.Name, p.Vehicle.Model)
	p.Vehicle.Honk(p.Name)
}

func main()  {
	car := Car{
		Model: "Mustang",
		Year: 2022,
		Manufacturer: "Ford",
	}

	person := Person{
		Name: "John",
		Age: 30,
		Vehicle: car,
	}

	person.Drive()
}