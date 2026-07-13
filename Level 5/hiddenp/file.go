package main

import (
	"fmt"
	"os"
)

func main() {
	// Must have exactly 2 arguments
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	// Empty s1 is always hidden
	if len(s1) == 0 {
		fmt.Println(1)
		return
	}

	i := 0 // index for s1

	for j := 0; j < len(s2) && i < len(s1); j++ {
		if s2[j] == s1[i] {
			i++
		}
	}

	if i == len(s1) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}