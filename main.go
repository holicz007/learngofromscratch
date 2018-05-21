package main

import (
	"fmt"
)

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

	/* here is simple project 1 */

	// cards := newDeck()
	// cards.suffler()
	// cards.print()

	/* here is simple project 2 */

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@exmalple.com",
			zipCode: 10250,
		},
	}

	jim.updateName("Jimmy")
	jim.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {

	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
