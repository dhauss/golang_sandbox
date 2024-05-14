package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	contact := contactInfo{email: "alex@email.com", zip: 12345}
	alex := person{"Alex", "Anderson", contact}
	fmt.Println(alex)

	alex2 := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact:   contact,
	}
	fmt.Println(alex2.firstName)

	var alex3 person
	fmt.Printf("%+v\n", alex3)
	alex3.firstName = "Alex"
	alex3.lastName = "Anderson"
	alex3.contact = contact
	alex3.print()

	alex3Pointer := &alex3
	alex3Pointer.updateName("Alexander")
	alex3.print()

	alex3.updateName("AlexAgain")
	alex3.print()

}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (pPointer *person) updateName(newFirstName string) {
	(*pPointer).firstName = newFirstName
}
