package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	axel := person{firstName: "Axel", lastName: "Morgan"}
	fmt.Println(axel.firstName, "", axel.lastName)
	axel.lastName = "Anderson"
	fmt.Println(axel)
}
