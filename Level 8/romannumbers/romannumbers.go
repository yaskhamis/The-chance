package group

import (
	"os"
	"fmt"
)

func StrConv(s string)int{
	result := 0
	sign := 1
	index := 0

	if s[0] == '-' {
		sign = -1
		index = 1
	}

	for i := index ; i < len(s) ; i++ {
		result = result*10 + int(s[i] - '0')
	}
	return result * sign
}

type Roman struct {
	Value int
	Symbol string
}

func Romana() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	num := StrConv(os.Args[1])
	if num < 1 || num > 3999 {
		fmt.Printf("ERROR: cannot convert to roman digit")
		return 
	}

	romans := []Roman{
		{1000, "M"},
		{900, "(M-C)"},
		{500, "D"},
		{400, "(D-C)"},
		{100, "C"},
		{90, "(C-X)"},
		{50, "L"},
		{40, "(L-X)"},
		{10, "X"},
		{9, "(X-I)"},
		{5, "V"},
		{4, "(V-I)"},
		{1, "I"},
	}

	// Build calculation string
	calc := ""
	val := num
	for _, r := range romans {
		for val >= r.Value {
			if len(calc) > 0 {
				calc += "+"
			}
			calc += r.Symbol
			val -= r.Value
		}
	}

	// Build compact Roman numeral string
	compact := ""
	val = num
	for _, r := range romans {
		for val >= r.Value {
			s := r.Symbol
			for i := len(s)-1 ; i >= 0; i-- {
				if s[i] != '(' && s[i] != ')' && s[i] != '-' {
					compact += string(s[i])
				}
			}
			val -= r.Value
		}
	}

	fmt.Println(calc)
	fmt.Println(compact)
}