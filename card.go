package main

import (
	"strings"
)

type card struct {
	suit  string
	value string
}

func buildCard(s string) card {

	sv := strings.Split(s, ":")
	suit := sv[0]
	value := sv[1]
	return card{suit: suit, value: value}
}

func (c card) print() {
	println(c.value + " of " + c.suit)
}

func (c card) toString() string {
	return c.suit + ":" + c.value
}
