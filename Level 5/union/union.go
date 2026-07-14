package main

import (
	"os"
)

func main() {
	if len(os.Args) != 3 {
		println()
		return
	}

	seen := make(map[rune]bool)

	for _, str := range os.Args[1:] {
		for _, ch := range str {
			if !seen[ch] {
				seen[ch] = true
				print(string(ch))
			}
		}
	}
	println()
}