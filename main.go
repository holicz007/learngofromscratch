package main

import "fmt"

func main() {

	card := newCard()

	fmt.Println(card)

}

/*
	การประกาศ fucntions ใน Go

	func functionName(parameters) returntype {


	}
*/

func newCard() string {
	return "Five of Diamonds"
}
