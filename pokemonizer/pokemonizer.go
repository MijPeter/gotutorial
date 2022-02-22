package pokemonizer

import (
	"math/rand"
	"strings"
	"unicode"
)

// is word consisting only of letters
func isWord(word string) bool {
	for _, letter := range word {
		if !unicode.IsLetter(letter) {
			return false
		}
	}
	return true
}

// randomly returns true or false
func randomBool() bool {
	return rand.Float32() < 0.5
}

// Pokemonizes input word and returns it.
func Pokemonize(input string) string {
	if !isWord(input) {
		return input
	}

	var output strings.Builder
	for _, letter := range input {
		if randomBool() {
			output.WriteRune(unicode.ToUpper(letter))
		} else {
			output.WriteRune(unicode.ToLower(letter))
		}
	}

	return output.String()
}
