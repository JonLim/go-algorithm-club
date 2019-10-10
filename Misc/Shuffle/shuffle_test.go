package shuffle

import (
	"reflect"
	"testing"
)

func TestDeckShuffle(t *testing.T) {
	tests := []struct {
		name string
		deck deck
	}{
		{
			name: "Deck with shuffle function returns a different order for cards",
			deck: deck{cards: []card{"10", "J", "Q", "K", "A"}},
		},
		{
			name: "Empty Deck with shuffle function returns empty deck",
			deck: deck{cards: []card{}},
		},
		{
			name: "Larger Deck with shuffle function returns a different order for cards",
			deck: deck{cards: []card{"10", "J", "Q", "K", "A", "2", "3", "4", "5", "6", "7", "8", "9"}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			originalDeckCards := append([]card{}, test.deck.cards)
			test.deck.shuffle()

			if reflect.DeepEqual(originalDeckCards, test.deck.cards) {
				t.Errorf("Expected deck with different order from %v, received deck %v instead.", originalDeckCards, test.deck.cards)
			}
		})
	}
}

func TestDeckFYKShuffle(t *testing.T) {
	tests := []struct {
		name string
		deck deck
	}{
		{
			name: "Deck with fykShuffle function returns a different order for cards",
			deck: deck{cards: []card{"10", "J", "Q", "K", "A"}},
		},
		{
			name: "Empty Deck with fykShuffle function returns empty deck",
			deck: deck{cards: []card{}},
		},
		{
			name: "Larger Deck with fykShuffle function returns a different order for cards",
			deck: deck{cards: []card{"10", "J", "Q", "K", "A", "2", "3", "4", "5", "6", "7", "8", "9"}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			originalDeckCards := append([]card{}, test.deck.cards)
			test.deck.fykShuffle()

			if reflect.DeepEqual(originalDeckCards, test.deck.cards) {
				t.Errorf("Expected deck with different order from %v, received deck %v instead.", originalDeckCards, test.deck.cards)
			}
		})
	}
}
