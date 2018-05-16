package main

import (
	"fmt"
)

func main() {

	cards := deck{"Ace of Diamonds", "Five of Diamonds"}

	fmt.Println(cards)

	cards.print()

}
