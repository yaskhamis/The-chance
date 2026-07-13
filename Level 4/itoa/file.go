package main

import "fmt"

func main() {
	fmt.Print(12/10)
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}

	result := ""

	for n > 0 {
		digit := n % 10
		result = string(rune(digit+'0')) + result
		n /= 10
	}

	return sign + result
}