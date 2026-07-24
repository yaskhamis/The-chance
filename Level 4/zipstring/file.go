package main

import (
	"fmt"
)

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}

func ZipString(s string) string {
	if s == "" {
		return ""
	}

	result := ""
	count := 1

	for i := 1; i <= len(s); i++ {
		if i < len(s) && s[i] == s[i-1] {
			count++
		} else {
			result += itoa(count)
			result += string(s[i-1])
			count = 1
		}
	}

	return result
}

func itoa(n int) string {
	result := ""
	for n > 0 {
		digit := n % 10
		result += string(rune(digit + '0'))
		n /= 10
	}
	return result
}
