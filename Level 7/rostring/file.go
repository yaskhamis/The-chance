package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}

	input := os.Args[1]

	// Split input into words manually (ignore extra spaces)
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

	if len(words) == 0 {
		fmt.Println()
		return
	}

	for i := 1; i < len(words); i++ {
		fmt.Print(words[i], " ")
	}
	fmt.Print(words[0])
	fmt.Println()
}
