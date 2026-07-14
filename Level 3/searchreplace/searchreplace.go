package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	if len(os.Args[2]) != 1 || len(os.Args[3]) != 1 {
		return
	}

	str := os.Args[1]
	old := rune(os.Args[2][0])
	new := rune(os.Args[3][0])

	for _, ch := range str {
		if ch == old {
			z01.PrintRune(new)
		} else {
			z01.PrintRune(ch)
		}
	}
	z01.PrintRune('\n')
}
