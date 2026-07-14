package piscine

func RepeatAlpha(s string) string {
	result := ""

	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			for i := 0; i < int(ch-'a')+1; i++ {
				result += string(ch)
			}
		} else if ch >= 'A' && ch <= 'Z' {
			for i := 0; i < int(ch-'A')+1; i++ {
				result += string(ch)
			}
		} else {
			result += string(ch)
		}
	}

	return result
}
