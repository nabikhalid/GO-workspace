package main

import "fmt"

// struct == object

type person struct {

	name string
	age int

}

func main() {

	CR7 := person{name: "Cristiano Ronaldo", age: 35}
	fmt.Println(CR7)
	fmt.Println(CR7.age)

}