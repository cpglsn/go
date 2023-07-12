package main

func main() {

	cards := newDeck()

	/*	hand := cards.deal(5)
		 	fmt.Println("Hand: ")
			hand.print()
			fmt.Println("Remaining cards: ")
			cards.print()
			fmt.Println(cards.toString())
	cards := newDeckFromFile("ciao") */

	//cards.print()

	cards.shuffle(2)

	cards.print()

}
