package group

import (
	"fmt"
	"os"
	// "strings"
)

func isVowel(ch byte) bool {
	ch = byte(ToLower(string(ch))[0])
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
}

func plglatin() {
	if len(os.Args) != 2 {
		// Do nothing if not exactly 1 argument
		return
	}

	word := os.Args[1]
	wordLower := ToLower(word)

	// Find first vowel
	index := -1
	for i := 0; i < len(wordLower); i++ {
		if isVowel(wordLower[i]) {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("No vowels")
		return
	}

	// Pig Latin transformation
	var pig string
	if index == 0 {
		// Starts with vowel
		pig = word + "ay"
	} else {
		// Starts with consonant(s)
		pig = word[index:] + word[:index] + "ay"
	}

	fmt.Println(pig)
}

func ToLower(s string) string{
	for i := 0 ; i < len(s) ; i++ {
		char := s[i]
		if char >= 'A' && char <= 'Z' {
			char = char + 32
		}
	}
	return s
}
