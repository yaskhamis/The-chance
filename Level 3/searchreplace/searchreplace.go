package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	Args := os.Args[1:]

	if len(Args) != 4 {
		return
	}
	str := []rune(Args[1])
	find := []rune(Args[2])
	replace := []rune(Args[3])

	for _, v := range str {
		if v == find[0] {
			z01.PrintRune(replace[0])
		} else {
			z01.PrintRune(v)
		}
	}
	z01.Print('\n')
}
