package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]

	if len(input) == 0 {
		fmt.Println()
		return
	}

	// Step 1: Split the string into words manually (without strings package)
	words := []string{}
	word := ""
	for i := 0; i < len(input); i++ {
		ch := input[i]
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

	// Step 2: Print words in reverse order
	for i := len(words) - 1; i >= 0; i-- {
		fmt.Print(words[i])
		if i != 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}