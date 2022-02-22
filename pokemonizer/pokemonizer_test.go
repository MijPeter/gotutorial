package pokemonizer_test

import (
	"example/hello/pokemonizer"
	"testing"
)

func TestPokemonizeSentence(t *testing.T) {
	// given
	word := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."

	// when
	pokemonizedWord := pokemonizer.Pokemonize(word)

	// then
	if word != pokemonizedWord {
		t.Fatal("Sentences differ.")
	}
}

func TestPokemonizeWord(t *testing.T) {
	// given
	word := "charizard"
	tries := 10 // for 10 or more chance of failing is (1 / 2^9)^10, in theory

	// when
	pokemonizedWord := word
	for tries > 0 && word == pokemonizedWord {
		pokemonizedWord = pokemonizer.Pokemonize(word)
		tries--
	}

	// then
	if word == pokemonizedWord {
		t.Fatal("Words don't differ.")
	}
}
