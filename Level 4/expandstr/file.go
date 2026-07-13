package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main(){
	CleanStr()
}

func CleanStr(){
	args := os.Args

	if len(args) != 2 {
		z01.PrintRune('\n')
	}
	args1 := args[1]


	words := []string{}
	word := ""
	for i := 0 ; i < len(args1); i++ {
		ch := args1[i]
		if string(ch) != " " {
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

	for i, char := range words {
		PrintStrr(char)
		if i != len(words) - 1 {
			z01.PrintRune(' ')
			z01.PrintRune(' ')
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune(' ')
}

func PrintStrr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}