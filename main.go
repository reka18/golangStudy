package main

func main() {
	deck := newDeckFromFile("deck.txt")
	deck.shuffle()
	deck.print()
}
