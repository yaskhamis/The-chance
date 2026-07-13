package main

import (
	"github.com/01-edu/z01"
	"fmt"
)

func main(){
	fmt.Println(CleanStr("Hello! How are you?"))
	fmt.Println(CleanStr("Hello How Are You"))
	fmt.Println(CleanStr("!!!!Whatsthis4"))
}

func CleanStr(s string) bool{

	words := []string{}
	word := ""
	for i := 0 ; i < len(s); i++ {
		ch := s[i]
		if (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') {
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

	for _, char := range words {
		for i := 0 ; i < len(char) ; i++ {
			ch := char[0]
			if ch >= 'a' && ch <= 'z' {
				return false
			}
		}
	}
	return true
	
}

func PrintStrr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}