package piscine

func CountAlpha(s string) int {
	c := 0
	for _, v := range s {
		if (v > 'a' && v < 'z') || (v > 'A' && v < 'Z') {
			c++
		}
	}
	return c
}
