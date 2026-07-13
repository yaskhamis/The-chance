package piscine

import "strings"

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}

	count := 0
	seen := ""

	for i := 0; i < len(str1); i++ {
		ch := string(str1[i])
		if !strings.Contains(str2, ch) && !strings.Contains(seen, ch) {
			seen += ch
			count++
		}
	}

	for i := 0; i < len(str2); i++ {
		ch := string(str2[i])
		if !strings.Contains(str1, ch) && !strings.Contains(seen, ch) {
			seen += ch
			count++
		}
	}

	return count
}
