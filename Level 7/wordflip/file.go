package main

import (
	"fmt"

)

func main() {
	fmt.Print(WordFlip("First second last"))
	fmt.Print(WordFlip(""))
	fmt.Print(WordFlip("     "))
	fmt.Print(WordFlip(" hello  all  of  you! "))
}

func WordFlip(s string) string {
	words := []string{}
	word := ""

	result := ""
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch != ' ' {
			word += string(ch)
		} else {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		}
	}
	if word != "" {
		words = append(words, word)
	}

	if len(words) == 0 {
		return "Invalid Output" + "\n"
	}

	for i := len(words)-1; i >= 0; i-- {
		result += words[i]
		if i != 0 {
			result += " "
		}
	}
	return result + "\n"
}
