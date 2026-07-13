package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main(){
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
	}

	args1 := os.Args[1]

	words := []string{}
	word := ""

	for i := 0 ; i < len(args1) ; i++ {
		char := args1[i]
		if char != ' ' {
			word += string(char)
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

	result := ""
	for _,word := range words {
		for i := 0 ; i < len(word) ; i++ {
			char := word[i]
			if i != len(word)-1 {
				result += string(ToLower(char))
			} else {
				result += string(ToUpper(char)) + " "
			}
		}
	}
	PrintStr(result)
}


func ToLower(char byte) byte {
	if char >= 'A' && char <= 'Z' {
		char += 32
	}
	return char
}

func ToUpper(char byte) byte {
	if char >= 'a' && char <= 'z' {
		char -= 32
	}
	return char
}

func PrintStr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}