package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 { // Checks if the argument count is not exactly 2 (plus program name)
		z01.PrintRune('\n')
		return
	}
	s1 := os.Args[1]               // first command-line argument to variable s1
	s2 := os.Args[2]               // second command-line argument to variable s2
	result := ""                   // Initializes an empty string to store common unique characters
	for i := 0; i < len(s1); i++ { // Loops through each byte of the first string s1
		alreadyPrinted := false            // Sets a flag to track if the current character is already processed
		for j := 0; j < len(result); j++ { // Loops through the accumulated result string
			if s1[i] == result[j] { // Compares current character with already stored characters
				alreadyPrinted = true
				break // Exits the inner loop early since duplicate is confirmed
			}
		}
		if alreadyPrinted { // Checks if the current character was marked as a duplicate
			continue // Skips the rest of the loop to process the next character
		}
		for k := 0; k < len(s2); k++ { // Loops through each byte of the second string s2
			if s1[i] == s2[k] { // Checks if the character from s1 exists in s2
				result += string(s1[i]) // Appends the matching character to the result string
				break                   // Exits the loop early to prevent multiple additions of the same match
			}
		}
	}
	for r := range result { // Iterates through the indices of the final result string
		z01.PrintRune(rune(result[r])) // Converts the byte at the current index to a rune and prints it
	} // Ends the final output loop
	z01.PrintRune('\n') // Prints a final trailing newline character
} // Ends the main function block
