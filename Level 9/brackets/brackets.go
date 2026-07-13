package group

import (
	"fmt"
)

func isCorrect(s string) bool {
	stack := []rune{}

	for _, ch := range s {
		// Opening brackets → push
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
		}

		// Closing brackets → check
		if ch == ')' || ch == ']' || ch == '}' {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if (ch == ')' && top != '(') ||
				(ch == ']' && top != '[') ||
				(ch == '}' && top != '{') {
				return false
			}
		}
	}

	// Stack must be empty at the end
	return len(stack) == 0
}

func Check(args []string){
	if len(args) < 2 {
		return
	}

	for _, arg := range args[1:] {
		if isCorrect(arg) {
			fmt.Println("OK")
		} else {
			fmt.Println("Error")
		}
	} 
}