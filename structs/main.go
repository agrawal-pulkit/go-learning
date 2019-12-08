package main

import "fmt"

type contactinfo struct {
	email string
	zip   int
}

type person struct {
	firstname string
	lastname  string
	contactinfo
}

func main() {
	contact := contactinfo{email: "agrawal.pulkit1994@gmail.com", zip: 203001}
	pulkit := person{firstname: "Pulkit", lastname: "Agrawal", contactinfo: contact}
	pulkit.print()

	var p1 person
	p1.firstname = "Aman"
	p1.lastname = "Agrawal"
	p1.contactinfo = contact
	p1.print()

	p2 := person{
		firstname: "test",
		lastname:  "test",
		contactinfo: contactinfo{
			email: "test@gmail.com",
			zip:   12345,
		},
	}
	p2.updateName("test2")
	p2.print()
}

//*person: means we are working with a pointer of person
//*pointerToPerson: This is an operator, means we want to manipulate the value the pointer is referencing
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstname = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
