package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	for i := 1; i < len(os.Args); i++ {
		str := os.Args[i]
		result := ""
		wordStart := -1

		for j := 0; j < len(str); j++ {
			c := str[j]

			// Check if character is a letter
			if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
				// Mark start of a word
				if wordStart == -1 {
					wordStart = j
				}
			} else {
				// Non-letter ends a word
				if wordStart != -1 {
					// Capitalize last letter of the word
					last := j - 1
					for k := wordStart; k < last; k++ {
						result += string(toLower(str[k]))
					}
					result += string(toUpper(str[last]))
					wordStart = -1
				}
				result += string(c)
			}
		}

		// Handle word ending at end of string
		if wordStart != -1 {
			last := len(str) - 1
			for k := wordStart; k < last; k++ {
				result += string(toLower(str[k]))
			}
			result += string(toUpper(str[last]))
		}

		fmt.Println(result)
	}
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

func toUpper(c byte) byte {
	if c >= 'a' && c <= 'z' {
		return c - 32
	}
	return c
}
