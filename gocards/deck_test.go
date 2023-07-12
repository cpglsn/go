package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expected := 16
	if len(d) != expected {
		t.Errorf("Expected deck lenght of %d, but got %d", expected, len(d))
	}
}

func TestSaveRestore(t *testing.T) {

	filename := "ciao_decktesting"
	os.Remove(filename)

	original := newDeck()

	tmp := make(deck, len(original))

	copy(tmp, original)

	//tmp.shuffle()
	tmp.saveToFile(filename)

	d := newDeckFromFile(filename)

	lo := len(original)
	ld := len(d)

	if lo != ld {
		t.Error("Different lenght of original and restored")
	}

	for i := 0; i < lo; i++ {
		if original[i] != d[i] {
			t.Error("Non matching elements")
			break
		}
	}

	os.Remove(filename)

}
