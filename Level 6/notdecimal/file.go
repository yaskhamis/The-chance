package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print(NotDecimal("0.1"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.1255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))
}

func NotDecimal(dec string) string {
	if dec == "" {
		return "\n"
	}
	sign := 1
	start := 0
	if dec[0] == '-' {
		start = 1
		sign = -1
	}

	// Find decimal point
	dotIndex := -1
	for i := start; i < len(dec); i++ {
		ch := dec[i]
		if ch == '.' {
			if dotIndex == -1 {
				dotIndex = i
			} else {
				// multiple dots → invalid
				return dec + "\n"
			}
		} else if ch < '0' || ch > '9' {
			// invalid char
			return dec + "\n"
		}
	}

	// No dot → return as-is
	if dotIndex == -1 {
		return dec + "\n"
	}

	// Count digits after decimal
	decimals := len(dec) - dotIndex - 1
	if decimals == 0 {
		return dec + "\n"
	}

	// Check if all digits after decimal are zeros
	allZero := true
	for i := dotIndex + 1; i < len(dec); i++ {
		if dec[i] != '0' {
			allZero = false
			break
		}
	}
	if allZero {
		return dec + "\n"
	}

	// Remove decimal point and convert to int by multiplying
	intPart := 0
	for i := 0; i < len(dec); i++ {
		if dec[i] != '.' {
			intPart = intPart*10 + int(dec[i]-'0')
		}
	}

	intPart *= sign

	return strconv.Itoa(intPart) + "\n"
}
