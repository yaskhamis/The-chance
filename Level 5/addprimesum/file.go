package main

import (
	"fmt"
	"os"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func atoi(s string) (int, bool) {
	if s == "" {
		return 0, false
	}

	num := 0
	start := 0
	negative := false

	if s[0] == '-' {
		negative = true
		start = 1
	}

	for i := start; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0, false
		}
		num = num*10 + int(c-'0')
	}

	if negative {
		num = -num
	}

	return num, true
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println(0)
		return
	}

	n, ok := atoi(os.Args[1])
	if !ok || n <= 0 {
		fmt.Println(0)
		return
	}

	sum := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			sum += i
		}
	}

	fmt.Println(sum)
}
