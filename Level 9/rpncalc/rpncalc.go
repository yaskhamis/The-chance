package group

import (
	"fmt"
	"strings"
	"strconv"
)

func Rpncalc(args []string){
	if len(args) < 2 {
		fmt.Println("Error")
		return
	}

	str := strings.Split(args[1], " ")
	str = cleanSplit(str)
	fmt.Println(str)
	result := 0
	for len(str) >= 3 {
		if Valid(str) {
			num1, _ := strconv.Atoi(str[0])
			num2, _ := strconv.Atoi(str[1])
			operator := str[2]
			str = str[3:]
			if operator == "-" {
				result = num1 - num2 
			} else if operator == "+" {
				result = num1 + num2
			} else if operator == "/" {
				result = num1 / num2 
			} else if operator == "%" {
				result = num1 % num2
			} else {
				result = num1 * num2
			}
			str = append([]string{convert(result)}, str...)
			fmt.Println(str)
		} else {
			fmt.Println("Error")
			return
		}
	}
	fmt.Println(result)
}

func Valid(s []string) bool {
	if len(s) < 3 {
		return false
	}

	_, err1 := strconv.Atoi(s[0])
	_, err2 := strconv.Atoi(s[1])
	if err1 != nil || err2 != nil {
		return false
	}

	if s[2] == "/" || s[2] == "%" || s[2] == "+" || s[2] == "-" || s[2] == "*" {
		return true
	}

	return false
}

func convert(num int) string {
	result := ""
	base := "0123456789"
	negative := false

	if num < 0 {
		negative = true
		num = -num
	}

	if num > 0 {
		remainder := num % 10
		result = string(base[remainder]) + result
		num /= 10
	}

	if negative {
		result = "-" + result
	}

	return result
}

func cleanSplit(s []string) []string {
	result := []string{}

	for _, v := range s {
		if v != "" {
			result = append(result, v)
		}
	}

	return result
}