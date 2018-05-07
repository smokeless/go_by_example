package main

import "fmt"

type person struct { //new struct with name and age fields
	name string
	age  int
}

func structGo() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 20})

	fmt.Println(person{name: "Fred"}) //missing fields are 0 valued

	fmt.Println(&person{name: "Ann", age: 80}) //pointer to struct

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s

	fmt.Println(sp.age)
	sp.age = 10
	//struct fields are mutable.
	fmt.Println(sp.age)
	fmt.Println(s)

}
