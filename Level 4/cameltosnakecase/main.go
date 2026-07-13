package main

import (
	"fmt"

)

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}


func CamelToSnakeCase(s string) string{
	array := []rune(s)
	result := s
	if len(s) == 0 {
		return ""
	}

	if IsCamleCase(s) {
		result = ""
		for i := 1; i < len(s) ; i++ {
			if i == 1 {
				result += string(array[0])
			}

			if array[i] >= 'A' && array[i] <= 'Z' {
				result += "_" + string(array[i])
			} else {
				result += string(array[i])
			}
		}
	} 
	return result
}

func IsCamleCase(s string)bool {
	array := []rune(s)
	for i := 0 ; i < len(s)-1; i++ {
		if !((array[i] >= 'a' && array[i] <= 'z') ||
			(array[i] >= 'A' && array[i] <= 'Z')) {
			return false
		}

		if array[i] >= 'A' && array[i] <= 'Z' &&
			array[i+1] >= 'A' && array[i+1] <= 'Z' {
			return false
		}
		
	}
	return true
}