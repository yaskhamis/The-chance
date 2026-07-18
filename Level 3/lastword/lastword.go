package piscine

func LastWord(s string) string {
	end := -1
	
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			end = i
			break
		}
	}
	if end == -1 {
		return "\n"
	}
	start := end
	for start >= 0 && s[start] != ' ' {
		start--
	}
	return s[start+1:end+1] + "\n"
}
