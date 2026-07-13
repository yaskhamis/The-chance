package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Must have exactly 1 argument
	if len(os.Args) != 2 {
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 1 {
		return
	}

	first := true
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			if !first {
				fmt.Print("*")
			}
			fmt.Print(i)
			first = false
			n /= i
		}
	}

	// If there's any remaining prime factor
	if n > 1 {
		if !first {
			fmt.Print("*")
		}
		fmt.Print(n)
	}

	fmt.Println()
}