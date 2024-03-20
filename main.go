package main

import (
	"flag"
	"math/rand"
	"strconv"

	petname "github.com/dustinkirkland/golang-petname"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	words   = flag.Int("words", 1, "The number of words in the pet name")
	numBots int
)

func main() {

	game := NewGame()

	println()
	bots := getInput("Number of bots : ")
	numBots, _ = strconv.Atoi(bots)

	for i := 0; i < numBots; i++ {
		var c ClassName
		if rand.Intn(2) == 0 {
			c = MAGE
		} else {
			c = WARRIOR
		}
		NewBot(c, cases.Title(language.English, cases.Compact).String(petname.Generate(*words, "-")), game)
	}

	game.Fight()

}
