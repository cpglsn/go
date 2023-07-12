package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := [4]string{"Spades", "Diamonds", "Hearts", "Clubs"}

	cardValues := [4]string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) deal(i int) deck {
	tmp := d[:i]
	d = d[i:]
	return tmp
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(raw), ","))
}

func (d deck) shuffle(numberOfShuffle ...int) {
	nos := 1
	if len(numberOfShuffle) > 0 {
		nos = numberOfShuffle[0]
	}

	lunghezza := len(d)

	for i := 1; i <= nos; i++ {

		for i := range d {
			randomPosition := rand.Intn(lunghezza)
			d[randomPosition], d[i] = d[i], d[randomPosition]
		}
	}
}
