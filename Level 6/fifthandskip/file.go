package main

import "fmt"

func main(){
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))
}

func FifthAndSkip(str string) string {
	if str == "" {
		return "\n"
	}


	nonSpaceCount := 0
	for _, char := range str {
		if char != ' ' {
			nonSpaceCount++
		}
	}

	if nonSpaceCount < 5 {
		return "Invalid Input\n"
	} 

	result := ""
	count := 0

	for i := 0 ; i < len(str); i++ {
		ch := str[i]

		if ch == ' ' {
			continue
		}

		count++
		if count%6 == 0 {
			continue
		}

		result += string(ch)
		if count % 5 == 0 {
			result += " "
		}
	}
	return result + "\n"
}