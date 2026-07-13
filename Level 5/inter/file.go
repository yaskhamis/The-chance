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

	result := ""

	for i := 0; i < len(s1); i++ {
		// Check if s1[i] was already printed
		alreadyPrinted := false
		for j := 0; j < len(result); j++ {
			if s1[i] == result[j] {
				alreadyPrinted = true
				break
			}
		}
		if alreadyPrinted {
			continue
		}

		// Check if s1[i] exists in s2
		for k := 0; k < len(s2); k++ {
			if s1[i] == s2[k] {
			result += string(s1[i])
			break
			}
		}
	}

	fmt.Println(result)
}