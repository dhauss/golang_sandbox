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
	fmt.Printf("%+v\n", alex3)

}
