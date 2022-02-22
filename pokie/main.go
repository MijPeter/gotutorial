package main

import (
	"example/hello/pokemonizer"
	"fmt"
	"math/rand"
	"time"
	"rsc.io/quote/v4"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(pokemonizer.Pokemonize(quote.Opt()))
}
