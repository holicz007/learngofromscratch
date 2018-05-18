package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Excepted deck lenght of 20, but got %d ", len(d))
	}
}
