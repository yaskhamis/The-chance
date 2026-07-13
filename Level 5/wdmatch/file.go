package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	i := 0 // index for s1

	for j := 0; j < len(s2) && i < len(s1); j++ {
		if s2[j] == s1[i] {
			i++
		}
	}

	if i == len(s1) {
		fmt.Println(s1)
	}
}