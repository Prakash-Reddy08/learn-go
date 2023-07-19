package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	//one way of assiginig values to struce
	// prakash := person{firstName: "Prakash",
	// 	lastName: "Mettu",
	// 	contact: contactInfo{
	// 		email:   "mjprakashreddy1@gmail.com",
	// 		zipCode: 13243,
	// 	},
	// }

	// another way
	var prakash person

	prakash.firstName = "Prakash"
	prakash.lastName = "Mettu"
	prakash.contact.email = "mjprakashredd1@gmail.com"
	prakash.contact.zipCode = 12324

	// if we try to update name using prakash.updateName(..)
	// by default in go the receiver function receives copy of prakash. it won't recive reference to prakash struct
	// so we use & operator to send the reference(i.e memory address) to the receiver
	(&prakash).updateName("Jyothi Prakash")
	//the above syntax can also be re-written like this
	// prakash.updateName("Jyothi Prakash")
	prakash.print()
}

// here * operator is used to access the value using the address
func (pointerToPerson *person) updateName(newFirstName string) {
	// now pointerToPerson contains the address of prakash
	// Go automatically handles the dereferencing for you, allowing you to directly access the firstName field using the pointerToPerson.firstName syntax.
	// This is because Go allows you to use the . operator on both pointers and non-pointers, automatically dereferencing the pointer when necessary.
	// so insted of this
	// (*pointerToPerson).firstName = newFirstName
	// we can write like this
	pointerToPerson.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
