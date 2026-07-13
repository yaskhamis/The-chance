package main

import "github.com/01-edu/z01"

func main() {
	first := true

	for a := '9'; a >= '2'; a-- {
		for b := a - 1; b >= '1'; b-- {
			for c := b - 1; c >= '0'; c-- {

				if !first {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}

				z01.PrintRune(a)
				z01.PrintRune(b)
				z01.PrintRune(c)

				first = false
			}
		}
	}

	z01.PrintRune('\n')
}