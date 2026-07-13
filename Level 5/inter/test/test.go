package main

import "fmt"

func main() {
	fmt.Print(Inter("padinton", "paqefwtdjetyiytjneytjoeyjnejeyj"))
}

func Inter(s1, s2 string)string {
	result := ""

	for i := 0; i < len(s1) ; i++ {
		isaleardyPrinted := false
		for j := 0 ; j < len(result) ; j++ {
			if s1[i] == result[j] {
				isaleardyPrinted = true
				break
			}
		}

		if !isaleardyPrinted {
			for h := 0 ; h < len(s2) ; h++ {
				if s1[i] == s2[h] {
					result += string(s1[i])
					break
				}	
			}
		}
	}
	return result
}
