package shuffle

import (
	"math/rand"
)

type card interface{}

type deck struct {
	cards []card
}

func (d *deck) shuffle() {
	if len(d.cards) == 0 {
		return
	}

	temp := []card{}

	var pluck func(cards []card)
	pluck = func(cards []card) {
		i := rand.Intn(len(cards))
		obj := cards[i]
		temp = append(temp, obj)
		cards = append(cards[:i], cards[i+1:]...) // Removes the "card"

		if len(cards) != 0 {
			pluck(cards) // Recursively pluck until `cards` are empty
		}
	}

	pluck(d.cards)
	d.cards = temp
}

// Fisher-Yates / Knuth shuffle function
func (d *deck) fykShuffle() {
	if len(d.cards) == 0 {
		return
	}

	for i := len(d.cards) - 1; i > 0; i-- {
		j := rand.Intn(i)

		if i != j {
			d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
		}
	}
}
