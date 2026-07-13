package main

import "github.com/01-edu/z01"

func main(){
	str:= " only    it's  harder   "
	PrintStr(FilterSpaces(str))
}

func FilterSpaces(s string) string {
	array := []rune(s)
	start := 0


	if array[0] == ' ' {
		start = CountSpaces(s, 0)
	}

	newStr := s[start:]
	result := newStr
	
	for i := 0 ; i < len(result) -1 ;{
		if result[i] == ' ' && result[i+1] == ' ' {
			result = result[:i] + result[i+1:]
		} else {
			i++
		}
	}

	for i := 0 ; i >= len(result) ;i++ {
		if len(result) > 0 && result[len(result)-1] == ' ' {
			return result
		} else {
			break
		}
	}
	return result
}

func CountSpaces(s string, start int) int{
	count := 0
	array := []rune(s)
	for i := start ; i < len(s); i++{
		if array[i] != ' ' {
			break
		} else {
			count++
		}
	}
	return count
}


func PrintStr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}