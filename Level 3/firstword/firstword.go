package piscine

func FirstWord(s string) string {
	start := -1
	for i, c := range s {
		if c != ' ' {
			start = i
			break
		}
	}
	if start == -1 {
		return "\n"
	}
	end := start
	for end < len(s) && s[end] != ' ' {
		end++
	}
	return s[start:end] + "\n"
}
