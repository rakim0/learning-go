package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string
	age  int
}

func main() {
	employee := struct {
		name string
		id   int
	}{
		name: "Alice",
		id:   121,
	}
	fmt.Println(reflect.TypeOf(employee))
	type Address struct {
		Street string
		City   string
	}
	type Contact struct {
		Name    string
		Address Address
		Phone   string
	}

	contact := Contact{
		Name: "Rakim",
		Address: Address{
			City:   "Prayagraj",
			Street: "Pipalgaon",
		},
		Phone: "7044550905",
	}

	fmt.Printf("%+v\n", contact)

	person := Person{
		name: "Rakim",
	}
	fmt.Println("Person's Old Name: ", person.name)
	person.modifyPerson("Melkey")
	fmt.Println("Person's New Name: ", person.name)

	x := 20
	x_ptr := &x
	fmt.Println("Value of x: ", x)
	*x_ptr = 1
	fmt.Println("New value of x:", x)
}

func (p *Person) modifyPerson(name string) {
	p.name = name
}
