# Shuffle

**Goal:** Rearrange the contents of an array.

Imagine you're making a card game and you need to shuffle a deck of cards. You can represent the deck by an array of `Card` objects and shuffling the deck means to change the order of those objects in the array. (It's like the opposite of sorting.)

Here is a naive way to approach this in Go:

```golang
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
```

To try it out, copy the code into a playground and then do:

```golang
list := deck{cards: []card{"a", "b", "c", "d", "e", "f", "g"}}
list.shuffle()
list.shuffle()
list.shuffle()
```

You should see three different arrangements -- or [permutations](../Combinatorics/) to use math-speak -- of the objects in the array.

This shuffle works *in place*, it modifies the contents of the original array. The algorithm works by creating a new array, `temp`, that is initially empty. Then we randomly choose an element from the original array and append it to `temp`, until the original array is empty. Finally, the temporary array is copied back into the original one.

This code works just fine but it's not very efficient. Removing an element from an array is an **O(n)** operation and we perform this **n** times, making the total algorithm **O(n^2)**. We can do better!

## The Fisher-Yates / Knuth shuffle

Here is a much improved version of the shuffle algorithm:

```golang
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
```

Again, this picks objects at random. In the naive version we placed those objects into a new temporary array so we could keep track of which objects were already shuffled and which still remained to be done. In this improved algorithm, however, we'll move the shuffled objects to the end of the original array. 

Let's walk through the example. We have the array:

	[ "a", "b", "c", "d", "e", "f", "g" ]

The loop starts at the end of the array and works its way back to the beginning. The very first random number can be any element from the entire array. Let's say it returns 2, the index of `"c"`. We swap `"c"` with `"g"` to move it to the end:

	[ "a", "b", "g", "d", "e", "f" | "c" ]
	             *                    *

The array now consists of two regions, indicated by the `|` bar. Everything to the right of the bar is shuffled already. 

The next random number is chosen from the range 0...6, so only from the region `[ "a", "b", "g", "d", "e", "f" ]`. It will never choose `"c"` since that object is done and we'll no longer touch it.

Let's say the random number generator picks 0, the index of `"a"`. Then we swap `"a"` with `"f"`, which is the last element in the unshuffled portion, and the array looks like this:

	[ "f", "b", "g", "d", "e" | "a", "c" ]
	   *                         *

The next random number is somewhere in `[ "f", "b", "g", "d", "e" ]`, so let's say it is 3. We swap `"d"` with `"e"`:

	[ "f", "b", "g", "e" | "d", "a", "c" ]
	                  *     *

And so on... This continues until there is only one element remaining in the left portion. For example:

	[ "b" | "e", "f", "g", "d", "a", "c" ]

There's nothing left to swap that `"b"` with, so we're done.

Because we only look at each array element once, this algorithm has a guaranteed running time of **O(n)**. It's as fast as you could hope to get!

## See also

These Swift implementations are based on pseudocode from the [Wikipedia article](https://en.wikipedia.org/wiki/Fisher–Yates_shuffle).

Mike Bostock has a [great visualization](http://bost.ocks.org/mike/shuffle/) of the shuffle algorithm.

*Reference material written for Swift Algorithm Club by Matthijs Hollemans*
*Adapted for Go Algorithm Club by [Jon Lim](https://github.com/JonLim)*