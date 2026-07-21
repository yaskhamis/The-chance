package piscine

func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}
	atWordStart := true
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == ' ' {
			atWordStart = true
			continue
		}
		if atWordStart {
			if c >= 'a' && c <= 'z' {
				return false
			}
			atWordStart = false
		}
	}
	return true
}
