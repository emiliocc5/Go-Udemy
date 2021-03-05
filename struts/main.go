package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	axel := person{firstName: "Axel", lastName: "Morgan"}
	fmt.Println(axel.firstName, "", axel.lastName)
	axel.lastName = "Anderson"
	fmt.Println(axel)

	var ezekiel person
	ezekiel.firstName = "Ezekiel"
	ezekiel.contact.email = "Test email"
	fmt.Println(ezekiel)

	jim := person{
		firstName: "jim",
		lastName:  "patterson",
		contact: contactInfo{
			email:   "fakeEmail123",
			zipCode: 123},
	}
	jim.updateName("jimmy")
	jim.print()

	//pointerToJim := &jim --> Not necessary
	//Go automatically turn your variable of type person into a pointer of that person
	jim.updateNameByPointer("jimmy")
	jim.print()
}

//Turn address into value with *address
//Turn value into address with &value

func (p *person) updateNameByPointer(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
